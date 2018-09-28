package db

import (
	"crypto-sync/litebit-api"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"strconv"
	"time"
)

func WriteToInfluxDB(apiResponse litebit_api.LitebitApiResult, databaseHost string, databasePort int) (map[string]interface{}, error) {
	buyRate, _ := strconv.ParseFloat(apiResponse.Result.Buy, 32)
	sellRate, _ := strconv.ParseFloat(apiResponse.Result.Sell, 32)

	connection, connectionErr := client.NewHTTPClient(client.HTTPConfig{
		Addr: fmt.Sprintf("http://%s:%d", databaseHost, databasePort),
	})
	if connectionErr != nil {
		return nil, connectionErr
	}
	defer connection.Close()
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "crypto",
		Precision: "s",
	})
	tags := map[string]string{"abbr": apiResponse.Result.Abbr, "name": apiResponse.Result.Name}
	fields := map[string]interface{}{
		"buy":  buyRate,
		"sell": sellRate,
	}
	pt, connectionErr := client.NewPoint("crypto", tags, fields, time.Now())
	if connectionErr != nil {
		return nil, connectionErr
	}
	bp.AddPoint(pt)
	writeErr := connection.Write(bp)
	if writeErr != nil {
		return nil, writeErr
	}

	return fields, nil
}
