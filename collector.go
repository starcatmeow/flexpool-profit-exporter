package main

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
)

type Collector struct{}

var (
	dailyProfitPerMHinCNY = prometheus.NewDesc("flexpool_daily_profit_per_mh_cny", "Daily Profit Per MH in CNY", []string{}, nil)
)

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- dailyProfitPerMHinCNY
}
func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	resp, err := http.Get("https://api.flexpool.io/v2/pool/coinsFull")
	if err != nil {
		log.Print("Error when fetching profit: ", err)
		return
	}
	var response Response
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}
	for _, coin := range response.Result {
		if coin.Ticker == "eth" {
			dailyETHRewardPerMH := coin.ChainData.DailyRewardPerGigaHashSec / 1000000000000000000000
			dailyCNYRewardPerMH := coin.MarketData.Prices.Cny * dailyETHRewardPerMH
			ch <- prometheus.MustNewConstMetric(dailyProfitPerMHinCNY, prometheus.GaugeValue, dailyCNYRewardPerMH)
			break
		}
	}
}
