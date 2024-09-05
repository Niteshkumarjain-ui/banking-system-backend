package main

import (
	"banking-system-backend/inbound"
	"banking-system-backend/util"
	"fmt"
)

func main() {
	if util.Configuration == nil {
		fmt.Println("Configuration not loaded properly")
	}

	util.GlobalWaitGroup.Add(1)
	inbound.HttpService()
	util.GlobalWaitGroup.Wait()
}
