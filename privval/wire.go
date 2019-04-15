package privval

import (
	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/XuanMaoSecLab/DolphinChain/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
	RegisterRemoteSignerMsg(cdc)
}
