package client_test

import (
	"os"
	"testing"

	"github.com/XuanMaoSecLab/DolphinChain/abci/example/kvstore"
	nm "github.com/XuanMaoSecLab/DolphinChain/node"
	rpctest "github.com/XuanMaoSecLab/DolphinChain/rpc/test"
)

var node *nm.Node

func TestMain(m *testing.M) {
	// start a tendermint node (and kvstore) in the background to test against
	app := kvstore.NewKVStoreApplication()
	node = rpctest.StartTendermint(app)

	code := m.Run()

	// and shut down proper at the end
	rpctest.StopTendermint(node)
	os.Exit(code)
}
