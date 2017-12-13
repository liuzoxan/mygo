package main

import log "github.com/cihub/seelog"

func main() {
	defer log.Flush()
	log.Trace("Hello from Seelog - Console!")
	log.Debug("Hello from Seelog - Console!")
	log.Info("Hello from Seelog - Console!")
	log.Warn("Hello from Seelog - Console!")
	log.Error("Hello from Seelog - Console!")
	log.Critical("Hello from Seelog - Console!")
	logger, err := log.LoggerFromConfigAsFile("./conf/seelog.xml")

	if err != nil {
		log.Error("Seelog load config failed!!")
	} else {
		log.ReplaceLogger(logger)
	}

	log.Info("Seelog using config")
	log.Trace("Hello from Seelog - Console!")
	log.Debug("Hello from Seelog - Console!")
	log.Info("Hello from Seelog - Console!")
	log.Warn("Hello from Seelog - Console!")
	log.Error("Hello from Seelog - Console!")
	log.Critical("Hello from Seelog - Console!")
}
