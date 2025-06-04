package scraping

import (
	"strconv"
	"strings"

	"regexp"

	"github.com/gocolly/colly/v2"
)

func NebiusScraper(collyAgent *colly.Collector) (float64, error) {
	c := collyAgent

	var priceHour string

	c.OnHTML("div.col.col-md-8.col-12.offset-md-1", func(e *colly.HTMLElement) {
		nebiusPriceHandler(e, &priceHour)
	})

	err := c.Visit("https://nebius.com/prices")
	if err != nil {
		return -1, err
	}

	priceHourFloat, err := strconv.ParseFloat(priceHour, 64)
	if err != nil {
		return -1, err
	}

	return priceHourFloat, nil
}

func nebiusPriceHandler(e *colly.HTMLElement, priceHour *string) {
	text := e.Text
	if strings.Contains(text, "NVIDIA H200 GPU") {
		re := regexp.MustCompile(`NVIDIA H200 GPU.*?\$(\d+\.\d+)`)
		match := re.FindStringSubmatch(text)
		if len(match) > 1 {
			*priceHour = match[1]
		}
	}
}
