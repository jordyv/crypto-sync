package main

import (
	"crypto-sync/db"
	"crypto-sync/litebit-api"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	var coins = flag.Args()
	if len(coins) == 0 {
		println("Usage: crypto-sync <coin code> <coin code>")
		return
	}

	fmt.Println("Start syncing...")
	for _, coin := range coins {
		response, err := litebit_api.GetMarket(coin)
		if err == nil && response.Success {
			if writtenFields, err := db.WriteToInfluxDB(response); err == nil {
				fmt.Println("Wrote", writtenFields)
			} else {
				println(err.Error())
			}
		} else {
			println(err.Error())
		}
	}
}
