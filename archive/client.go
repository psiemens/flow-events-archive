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

package archive

import (
	"context"

	"github.com/onflow/flow-go/engine/common/rpc/convert"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow/protobuf/go/flow/access"
	legacyaccess "github.com/onflow/flow/protobuf/go/flow/legacy/access"
	legacyentities "github.com/onflow/flow/protobuf/go/flow/legacy/entities"
	"google.golang.org/grpc"
)

type AccessClient interface {
	GetEventsForHeightRange(
		ctx context.Context, eventType string, startHeight, endHeight uint64,
	) ([]flow.BlockEvents, error)

	Close() error
}

type Client struct {
	rpcClient access.AccessAPIClient
	close     func() error
}

func NewClient(addr string, opts ...grpc.DialOption) (*Client, error) {
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	grpcClient := access.NewAccessAPIClient(conn)

	return &Client{
		rpcClient: grpcClient,
		close:     func() error { return conn.Close() },
	}, nil
}

func (c *Client) GetEventsForHeightRange(
	ctx context.Context,
	eventType string,
	startHeight,
	endHeight uint64,
) ([]flow.BlockEvents, error) {
	req := access.GetEventsForHeightRangeRequest{
		Type:        eventType,
		StartHeight: startHeight,
		EndHeight:   endHeight,
	}

	res, err := c.rpcClient.GetEventsForHeightRange(ctx, &req)
	if err != nil {
		return nil, err
	}

	results := res.GetResults()

	blockResults := make([]flow.BlockEvents, len(results))

	for i, result := range results {
		blockResults[i] = flow.BlockEvents{
			BlockID:        convert.MessageToIdentifier(result.BlockId),
			BlockHeight:    result.GetBlockHeight(),
			BlockTimestamp: result.GetBlockTimestamp().AsTime(),
			Events:         convert.MessagesToEvents(result.Events),
		}
	}

	return blockResults, nil
}

func (c *Client) Close() error {
	return c.close()
}

type LegacyClient struct {
	rpcClient legacyaccess.AccessAPIClient
	close     func() error
}

func NewLegacyClient(addr string, opts ...grpc.DialOption) (*LegacyClient, error) {
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	grpcClient := legacyaccess.NewAccessAPIClient(conn)

	return &LegacyClient{
		rpcClient: grpcClient,
		close:     func() error { return conn.Close() },
	}, nil
}

func (c *LegacyClient) GetEventsForHeightRange(
	ctx context.Context,
	eventType string,
	startHeight,
	endHeight uint64,
) ([]flow.BlockEvents, error) {
	req := legacyaccess.GetEventsForHeightRangeRequest{
		Type:        eventType,
		StartHeight: startHeight,
		EndHeight:   endHeight,
	}

	res, err := c.rpcClient.GetEventsForHeightRange(ctx, &req)
	if err != nil {
		return nil, err
	}

	results := res.GetResults()

	blockResults := make([]flow.BlockEvents, len(results))

	for i, result := range results {
		blockResults[i] = flow.BlockEvents{
			BlockID:     convert.MessageToIdentifier(result.BlockId),
			BlockHeight: result.GetBlockHeight(),
			Events:      legacyMessagesToEvents(result.Events),
		}
	}

	return blockResults, nil
}

func legacyMessagesToEvents(m []*legacyentities.Event) []flow.Event {
	events := make([]flow.Event, len(m))

	for i, event := range m {
		events[i] = legacyMessageToEvent(event)
	}

	return events
}

func legacyMessageToEvent(m *legacyentities.Event) flow.Event {
	return flow.Event{
		Type:             flow.EventType(m.GetType()),
		TransactionID:    flow.HashToID(m.GetTransactionId()),
		TransactionIndex: m.GetTransactionIndex(),
		EventIndex:       m.GetEventIndex(),
		Payload:          m.GetPayload(),
	}
}

func (c *LegacyClient) Close() error {
	return c.close()
}
