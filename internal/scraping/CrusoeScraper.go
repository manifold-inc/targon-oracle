package scraping

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func CrusoeScraper() float64 {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	var priceHour string

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		rowText := e.Text
		if strings.Contains(strings.ToLower(rowText), "h200") {
			priceHour = strings.Split(rowText, "$")[1]
			priceHour = strings.Fields(priceHour)[0]
		}
	})

	err := c.Visit("https://crusoe.ai/cloud/pricing?instance=3&numOfGpus=0")
	if err != nil {
		log.Fatal(err)
	}

	priceHourFloat, err := strconv.ParseFloat(priceHour, 64)
	if err != nil {
		log.Fatal(err)
	}

	return priceHourFloat
}