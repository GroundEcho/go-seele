/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */

package light

import (
	"github.com/seeleteam/go-seele/common"
	"github.com/seeleteam/go-seele/core/store"
	"github.com/seeleteam/go-seele/core/types"
)

type odrBlock struct {
	OdrItem
	Hash   common.Hash  // Retrieved block hash
	Height uint64       // Retrieved block height
	Block  *types.Block `rlp:"nil"` // Retrieved block
}

func (odr *odrBlock) code() uint16 {
	return blockRequestCode
}

func (odr *odrBlock) handle(lp *LightProtocol) (uint16, odrResponse) {
	var err error

	if odr.Hash.IsEmpty() {
		if odr.Block, err = lp.chain.GetStore().GetBlockByHeight(odr.Height); err != nil {
			lp.log.Debug("Failed to get block, height = %d, error = %v", odr.Height, err)
			odr.Error = err.Error()
		}
	} else if odr.Block, err = lp.chain.GetStore().GetBlock(odr.Hash); err != nil {
		lp.log.Debug("Failed to get block, hash = %v, error = %v", odr.Hash, err)
		odr.Error = err.Error()
	}

	return blockResponseCode, odr
}

func (odr *odrBlock) validate(request odrRequest, bcStore store.BlockchainStore) error {
	if odr.Block == nil {
		return nil
	}

	var err error

	if err = odr.Block.Validate(); err != nil {
		return err
	}

	hash := request.(*odrBlock).Hash
	if hash.IsEmpty() {
		if hash, err = bcStore.GetBlockHash(odr.Height); err != nil {
			return err
		}
	}

	if !hash.Equal(odr.Block.HeaderHash) {
		return types.ErrBlockHashMismatch
	}

	return nil
}
