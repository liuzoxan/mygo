package util

import (
	"github.com/cihub/seelog"
)

type Log struct {
	Runtime   seelog.LoggerInterface
	StatusLog seelog.LoggerInterface
}

func (log *Log) Init(runtime, status string) error {
	var err error

	if log.Runtime, err = seelog.LoggerFromConfigAsFile(runtime); err != nil {
		return err
	}

	if log.StatusLog, err = seelog.LoggerFromConfigAsFile(status); err != nil {
		return err
	}
	return nil
}

func (log *Log) Flush() {
	log.Runtime.Flush()
	log.StatusLog.Flush()
}

var Logger *Log

func init() { // init automatically before main
	Logger = new(Log)
}
