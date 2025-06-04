package scraping

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func TogetherScraper() float64 {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	var priceHour string

	c.OnHTML("li.pricing_content-row", func(e *colly.HTMLElement) {
		text := e.Text
		if strings.Contains(text, "H200 141GB") {
			priceHour = strings.Split(text, "$")[2]
		}
	})

	err := c.Visit("https://www.together.ai/pricing")
	if err != nil {
		log.Fatal(err)
	}

	priceHourFloat, err := strconv.ParseFloat(priceHour, 64)
	if err != nil {
		log.Fatal(err)
	}

	return priceHourFloat
}