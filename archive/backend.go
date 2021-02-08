/*
 * Flow Emulator
 *
 * Copyright 2019-2020 Dapper Labs, Inc.
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
	"fmt"

	"github.com/onflow/flow-go/access"
	"github.com/onflow/flow-go/model/flow"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Backend struct {
	log   zerolog.Logger
	nodes []*AccessNode
}

func NewBackend(log zerolog.Logger, nodes []*AccessNode) *Backend {
	return &Backend{
		log:   log,
		nodes: nodes,
	}
}

func (b *Backend) getNodeForHeightRange(startHeight, endHeight uint64) (*AccessNode, error) {
	for _, node := range b.nodes {
		if startHeight <= node.endHeight && endHeight <= node.endHeight {
			return node, nil
		}
	}

	return nil, fmt.Errorf("no valid archive node for height range (%d, %d)", startHeight, endHeight)
}

func (b *Backend) Ping(ctx context.Context) error {
	panic("implement me")
}

func (b *Backend) GetNetworkParameters(ctx context.Context) access.NetworkParameters {
	panic("implement me")
}

func (b *Backend) GetLatestBlockHeader(ctx context.Context, isSealed bool) (*flow.Header, error) {
	panic("implement me")
}

func (b *Backend) GetBlockHeaderByHeight(ctx context.Context, height uint64) (*flow.Header, error) {

	panic("implement me")
}

func (b *Backend) GetBlockHeaderByID(ctx context.Context, id flow.Identifier) (*flow.Header, error) {
	panic("implement me")
}

func (b *Backend) GetLatestBlock(ctx context.Context, isSealed bool) (*flow.Block, error) {
	panic("implement me")
}

func (b *Backend) GetBlockByHeight(ctx context.Context, height uint64) (*flow.Block, error) {
	panic("implement me")
}

func (b *Backend) GetBlockByID(ctx context.Context, id flow.Identifier) (*flow.Block, error) {
	panic("implement me")
}

func (b *Backend) GetCollectionByID(ctx context.Context, id flow.Identifier) (*flow.LightCollection, error) {
	panic("implement me")
}

func (b *Backend) SendTransaction(ctx context.Context, tx *flow.TransactionBody) error {
	panic("implement me")
}

func (b *Backend) GetTransaction(ctx context.Context, id flow.Identifier) (*flow.TransactionBody, error) {
	panic("implement me")
}

func (b *Backend) GetTransactionResult(ctx context.Context, id flow.Identifier) (*access.TransactionResult, error) {
	panic("implement me")
}

func (b *Backend) GetAccount(ctx context.Context, address flow.Address) (*flow.Account, error) {
	panic("implement me")
}

func (b *Backend) GetAccountAtLatestBlock(ctx context.Context, address flow.Address) (*flow.Account, error) {
	panic("implement me")
}

func (b *Backend) GetAccountAtBlockHeight(ctx context.Context, address flow.Address, height uint64) (*flow.Account, error) {
	panic("implement me")
}

func (b *Backend) ExecuteScriptAtLatestBlock(ctx context.Context, script []byte, arguments [][]byte) ([]byte, error) {
	panic("implement me")
}

func (b *Backend) ExecuteScriptAtBlockHeight(ctx context.Context, blockHeight uint64, script []byte, arguments [][]byte) ([]byte, error) {
	panic("implement me")
}

func (b *Backend) ExecuteScriptAtBlockID(ctx context.Context, blockID flow.Identifier, script []byte, arguments [][]byte) ([]byte, error) {
	panic("implement me")
}

func (b *Backend) GetEventsForHeightRange(
	ctx context.Context,
	eventType string,
	startHeight,
	endHeight uint64,
) ([]flow.BlockEvents, error) {
	node, err := b.getNodeForHeightRange(startHeight, endHeight)
	if err != nil {
		log.Err(err).Msg("failed to find archive node")
		return nil, err
	}

	results, err := node.client.GetEventsForHeightRange(ctx, eventType, startHeight, endHeight)
	if err != nil {
		log.Err(err).
			Str("eventType", eventType).
			Uint64("startHeight", startHeight).
			Uint64("endHeight", endHeight).
			Msg("failed get events from archive node")
		return nil, err
	}

	return results, nil
}

func (b *Backend) GetEventsForBlockIDs(ctx context.Context, eventType string, blockIDs []flow.Identifier) ([]flow.BlockEvents, error) {
	panic("implement me")
}
