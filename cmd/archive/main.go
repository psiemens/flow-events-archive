/*
 * Flow Events Archive
 *
 * Copyright 2019-2021 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"os"

	"github.com/onflow/flow-go/model/flow"
	"github.com/psiemens/graceland"
	"github.com/psiemens/sconfig"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/psiemens/flow-events-archive/archive"
)

const envPrefix = "FLOW"

type Config struct {
	GRPCPort  int  `default:"9000" info:"port to run gRPC server"`
	HTTPPort  int  `default:"8080" info:"port to run HTTP server"`
	GRPCDebug bool `default:"false" info:"enable gRPC debug logs"`
}

var conf Config

var (
	defaultHTTPHeaders = []archive.HTTPHeader{
		{
			Key:   "Access-Control-Allow-Origin",
			Value: "*",
		},
		{
			Key:   "Access-Control-Allow-Methods",
			Value: "POST, GET, OPTIONS, PUT, DELETE",
		},
		{
			Key:   "Access-Control-Allow-Headers",
			Value: "*",
		},
	}
)

func init() {
	err := sconfig.New(&conf).
		FromEnvironment(envPrefix).
		Parse()
	if err != nil {
		panic(err)
	}
}

func main() {
	log := zerolog.New(os.Stderr)

	backend := archive.NewBackend(
		log,
		[]*archive.AccessNode{
			archive.NewNode(
				1065711, 2033591,
				newLegacyClient("access-001.candidate4.nodes.onflow.org:9000"),
			),
			archive.NewNode(
				2033592, 3187930,
				newLegacyClient("access-001.candidate5.nodes.onflow.org:9000"),
			),
			archive.NewNode(
				3187931, 4132132,
				newLegacyClient("access-001.candidate6.nodes.onflow.org:9000"),
			),
			archive.NewNode(
				4132133, 4972986,
				newLegacyClient("access-001.candidate7.nodes.onflow.org:9000"),
			),
			archive.NewNode(
				4972987, 6483245,
				newLegacyClient("access-001.candidate8.nodes.onflow.org:9000"),
			),
			archive.NewNode(
				6483246, 7601062,
				newLegacyClient("access-001.candidate9.nodes.onflow.org:9000"),
			),
			archive.NewNode(
				7601063, 8742958,
				newClient("access-001.mainnet1.nodes.onflow.org:9000"),
			),
			archive.NewNode(
				8742959, 9737132,
				newClient("access-001.mainnet2.nodes.onflow.org:9000"),
			),
			archive.NewNode(
				9737133, 9992019,
				newClient("access-001.mainnet3.nodes.onflow.org:9000"),
			),
			archive.NewNode(
				9992020, 12020336,
				newClient("access-001.mainnet4.nodes.onflow.org:9000"),
			),
		},
	)

	chain := flow.Mainnet.Chain()

	grpcServer := archive.NewGRPCServer(backend, conf.GRPCPort, chain, conf.GRPCDebug)
	httpServer := archive.NewHTTPServer(grpcServer, conf.HTTPPort, defaultHTTPHeaders)

	group := graceland.NewGroup()

	group.Add(grpcServer)
	group.Add(httpServer)

	log.Info().Msgf("Listening for gRPC traffic on port %d...", conf.GRPCPort)
	log.Info().Msgf("Listening for HTTP traffic on port %d...", conf.HTTPPort)

	err := group.Start()
	if err != nil {
		log.Err(err).Msg("shutting down due to fatal error")
	}
}

func newClient(host string) archive.AccessClient {
	c, err := archive.NewClient(host, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024 * 1024 * 12))) // default of 4194304 is too small for some dense blocks (e.g. 11172812)
	if err != nil {
		panic(err)
	}

	return c
}

func newLegacyClient(host string) archive.AccessClient {
	c, err := archive.NewLegacyClient(host, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return c
}
