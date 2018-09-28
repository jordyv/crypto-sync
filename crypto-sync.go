package main

import (
	"crypto-sync/db"
	"crypto-sync/litebit-api"
	"flag"
	"fmt"
	"os"
	"time"
)

func SyncCoins(coins []string, databaseHost string, databasePort int) {
	fmt.Println("Start syncing...")
	for _, coin := range coins {
		response, err := litebit_api.GetMarket(coin)
		if err == nil && response.Success {
			if writtenFields, err := db.WriteToInfluxDB(response, databaseHost, databasePort); err == nil {
				fmt.Println("Wrote", writtenFields)
			} else {
				println(err.Error())
			}
		} else {
			println(err.Error())
		}
	}
}

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: crypto-sync [<flags>] <coin code> <coin code>")
		flag.PrintDefaults()
	}
	var timeBetweenSyncs = flag.Duration("time", time.Second*10, "Time between syncs, eg 10s or 5m")
	var databaseHost = flag.String("host", "localhost", "Host for InfluxDB")
	var databasePort = flag.Int("port", 6086, "Port for InfluxDB")
	flag.Parse()

	var coins = flag.Args()
	if len(coins) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	fmt.Println("Sync every", timeBetweenSyncs)

	for {
		go SyncCoins(coins, *databaseHost, *databasePort)
		time.Sleep(*timeBetweenSyncs)
	}
	os.Exit(0)
}
