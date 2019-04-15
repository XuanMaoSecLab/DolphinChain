package proxy

import (
	cmn "github.com/XuanMaoSecLab/DolphinChain/libs/common"
	dbm "github.com/XuanMaoSecLab/DolphinChain/libs/db"
	log "github.com/XuanMaoSecLab/DolphinChain/libs/log"
	"github.com/XuanMaoSecLab/DolphinChain/lite"
	lclient "github.com/XuanMaoSecLab/DolphinChain/lite/client"
)

func NewVerifier(chainID, rootDir string, client lclient.SignStatusClient, logger log.Logger, cacheSize int) (*lite.DynamicVerifier, error) {

	logger = logger.With("module", "lite/proxy")
	logger.Info("lite/proxy/NewVerifier()...", "chainID", chainID, "rootDir", rootDir, "client", client)

	memProvider := lite.NewDBProvider("trusted.mem", dbm.NewMemDB()).SetLimit(cacheSize)
	lvlProvider := lite.NewDBProvider("trusted.lvl", dbm.NewDB("trust-base", dbm.LevelDBBackend, rootDir))
	trust := lite.NewMultiProvider(
		memProvider,
		lvlProvider,
	)
	source := lclient.NewProvider(chainID, client)
	cert := lite.NewDynamicVerifier(chainID, trust, source)
	cert.SetLogger(logger) // Sets logger recursively.

	// TODO: Make this more secure, e.g. make it interactive in the console?
	_, err := trust.LatestFullCommit(chainID, 1, 1<<63-1)
	if err != nil {
		logger.Info("lite/proxy/NewVerifier found no trusted full commit, initializing from source from height 1...")
		fc, err := source.LatestFullCommit(chainID, 1, 1)
		if err != nil {
			return nil, cmn.ErrorWrap(err, "fetching source full commit @ height 1")
		}
		err = trust.SaveFullCommit(fc)
		if err != nil {
			return nil, cmn.ErrorWrap(err, "saving full commit to trusted")
		}
	}

	return cert, nil
}
