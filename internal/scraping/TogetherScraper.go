package scraping

import (
	"strconv"
	"strings"
	"github.com/gocolly/colly/v2"
)

func TogetherScraper(collyAgent *colly.Collector) (float64, error) {
	c := collyAgent

	var priceHour string

	c.OnHTML("li.pricing_content-row", func(e *colly.HTMLElement) {
		togetherPriceHandler(e, &priceHour)
	})

	err := c.Visit("https://www.together.ai/pricing")
	if err != nil {
		return -1, err
	}

	priceHourFloat, err := strconv.ParseFloat(priceHour, 64)
	if err != nil {
		return -1, err
	}

	return priceHourFloat, nil
}

func togetherPriceHandler(e *colly.HTMLElement, priceHour *string) {
	text := e.Text
	if strings.Contains(text, "H200 141GB") {
		*priceHour = strings.Split(text, "$")[2]
	}
}