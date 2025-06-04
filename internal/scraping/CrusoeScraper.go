package scraping

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func CrusoeScraper(collyAgent *colly.Collector) (float64, error) {
	c := collyAgent

	var priceHour string

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		crusoePriceHandler(e, &priceHour)
	})

	err := c.Visit("https://crusoe.ai/cloud/pricing?instance=3&numOfGpus=0")
	if err != nil {
		return -1, err
	}

	priceHourFloat, err := strconv.ParseFloat(priceHour, 64)
	if err != nil {
		return -1, err
	}

	return priceHourFloat, nil
}

func crusoePriceHandler(e *colly.HTMLElement, priceHour *string) {
	rowText := e.Text
	if strings.Contains(strings.ToLower(rowText), "h200") {
		*priceHour = strings.Split(rowText, "$")[1]
		*priceHour = strings.Fields(*priceHour)[0]
	}
}
