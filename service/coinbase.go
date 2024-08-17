package service

import (
	"encoding/json"
	"goro/model"
	"log"
	"net/http"
)

func GetAssets() ([]model.AssetIn, error) {
	apiUrl := GetContext().CoinbaseApiUrl
	resp, err := http.Get(apiUrl)

	if err != nil {
		return nil, err
	}

	var jsonData []model.AssetIn
	err = json.NewDecoder(resp.Body).Decode(&jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil

}

func GetRates(exchangeUrl, currency string) (any, error) {
	log.Println(exchangeUrl)
	req, _ := http.NewRequest("GET", exchangeUrl, nil)
	q := req.URL.Query()
	q.Add("currency", currency)
	req.URL.RawQuery = q.Encode()
	resp, err := (&http.Client{}).Do(req)

	if err != nil {
		log.Fatal("could not fetch rates:", err)
		return nil, err
	}

	var jsonData any
	err = json.NewDecoder(resp.Body).Decode(&jsonData)

	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
