package blockchain

import (
	amino "github.com/tendermint/go-amino"
	"github.com/XuanMaoSecLab/DolphinChain/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterBlockchainMessages(cdc)
	types.RegisterBlockAmino(cdc)
}
