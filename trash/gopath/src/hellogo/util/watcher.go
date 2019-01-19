package util

import (
	"bytes"
	"strconv"
	"sync"
	"time"
)

type watch struct {
	mutex   *sync.Mutex
	content map[string]int64
}

var watcher watch

func AddWatchValue(key string, value int64) {
	watcher.mutex.Lock()
	defer watcher.mutex.Unlock()
	v, in := watcher.content[key]
	if in {
		v = v + value
	} else {
		v = value
	}

	watcher.content[key] = v
	return
}

func SetWatchValue(key string, value int64) {
	watcher.mutex.Lock()
	defer watcher.mutex.Unlock()
	watcher.content[key] = value
}

func RunWatch() {
	go func() {
		for {
			select {
			case <-time.Tick(5 * time.Second):
				WriteDisk()
			}
		}
	}()
}

func WriteDisk() {
	var buffer bytes.Buffer
	watcher.mutex.Lock()
	defer watcher.mutex.Unlock()

	for key, value := range watcher.content {
		buffer.WriteString("\t")
		buffer.WriteString(key)
		buffer.WriteString(":")
		buffer.WriteString(strconv.FormatInt(value, 10))
		watcher.content[key] = 0
	}

	// NOTE: Please init logger before using
	Logger.StatusLog.Info(buffer.String())
}

func init() {
	watcher.mutex = new(sync.Mutex)
	watcher.content = make(map[string]int64)
}
