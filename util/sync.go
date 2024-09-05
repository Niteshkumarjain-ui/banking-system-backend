package util

import "sync"

var GlobalWaitGroup *sync.WaitGroup

func init() {
	GlobalWaitGroup = initGlobalWaitGroup()
}

func initGlobalWaitGroup() (globalWaitGroup *sync.WaitGroup) {
	globalWaitGroup = new(sync.WaitGroup)
	return
}
