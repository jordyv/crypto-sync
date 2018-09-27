package litebit_api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type LitebitApiMarketResult struct {
	Name      string `json:"name"`
	Abbr      string `json:"abbr"`
	Available string `json:"available"`
	Volume    string `json:"volume"`
	Buy       string `json:"buy"`
	Sell      string `json:"sell"`
}

type LitebitApiResult struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Result  LitebitApiMarketResult `json:"result"`
}

func GetMarket(coin string) (LitebitApiResult, error) {
	response, err := http.Get("https://api.litebit.eu/market/" + coin)
	if err != nil {
		return LitebitApiResult{}, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return LitebitApiResult{}, err
	}
	result := LitebitApiResult{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		return LitebitApiResult{}, jsonErr
	}

	return result, nil
}
