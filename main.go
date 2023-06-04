package main

import (
	"horus-watcher/configs"
	"horus-watcher/util"
	"time"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			util.ManageServices()
		}
	}
}
