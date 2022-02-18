package main

type Coin struct {
	Ticker    string `json:"ticker"`
	ChainData struct {
		DailyRewardPerGigaHashSec float64 `json:"dailyRewardPerGigaHashSec"`
	} `json:"chainData"`
	MarketData struct {
		Prices struct {
			Cny float64 `json:"cny"`
			Usd float64 `json:"usd"`
		} `json:"prices"`
	} `json:"marketData"`
}

type Response struct {
	Result []Coin `json:"result"`
}
