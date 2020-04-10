package testutils

import (
	log "github.com/ChainSafe/log15"
)

func SetLogger(lvl log.Lvl) {
	logger := log.Root()
	handler := logger.GetHandler()
	log.Root().SetHandler(log.LvlFilterHandler(lvl, handler))
}
