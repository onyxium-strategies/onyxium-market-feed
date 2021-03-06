package main

// TODO In the future this needs to change to a bittrex api wrapper
// TODO Exchange interface
import (
	"encoding/json"
	"github.com/onyxium-strategies/onyxium-strategy-worker/models"
	"log"
	"net/http"
)

type MarketSummaryResponse struct {
	Success       bool            `json:"success"`
	Message       string          `json:"message"`
	MarketSummary []models.Market `json:"result"`
}

/*
url: https://bittrex.com/api/v1.1/public/getmarketsummary?market=btc-ltc
response:
{   "success":true,
    "message":"",
    "result":[{
        "MarketName":"BTC-LTC",
        "High":0.02129998,
        "Low":0.01951000,
        "Volume":56074.33572962,
        "Last":0.02019990,
        "BaseVolume":1138.16360733,
        "TimeStamp":"2018-02-21T20:07:38.94",
        "Bid":0.02011701,
        "Ask":0.02019986,
        "OpenBuyOrders":2126,
        "OpenSellOrders":4863,
        "PrevDay":0.02064007,
        "Created":"2014-02-13T00:00:00"
    }]
}
*/

// Return
func getMarketSummary() (map[string]models.Market, error) {
	url := "https://bittrex.com/api/v1.1/public/getmarketsummaries"

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return make(map[string]models.Market), err
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return make(map[string]models.Market), err
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record MarketSummaryResponse

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	market := make(map[string]models.Market)

	for _, i := range record.MarketSummary {
		market[i.MarketName] = i
	}
	return market, nil
}
