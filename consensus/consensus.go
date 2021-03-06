/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */

package consensus

import (
	"github.com/seeleteam/go-seele/core/store"
	"github.com/seeleteam/go-seele/core/types"
	"github.com/seeleteam/go-seele/rpc"
)

type Engine interface {
	// Prepare header before generate block
	Prepare(store store.BlockchainStore, header *types.BlockHeader) error

	// VerifyHeader verify block header
	VerifyHeader(store store.BlockchainStore, header *types.BlockHeader) error

	// Seal generate block
	Seal(store store.BlockchainStore, block *types.Block, stop <-chan struct{}, results chan<- *types.Block) error

	// APIs returns the RPC APIs this consensus engine provides.
	APIs() []rpc.API

	// SetThreads set miner threads
	SetThreads(thread int)
}
