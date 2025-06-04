package scraping

import (
	"strconv"
	"strings"
	"github.com/gocolly/colly/v2"
)

func VastScraper(collyAgent *colly.Collector) (float64, error) {
	c := collyAgent

	var priceHour string

	c.OnHTML("tbody.css-yp9msd.elmp00v6 tr", func(e *colly.HTMLElement) {
		vastPriceHandler(e, &priceHour)
	})

	err := c.Visit("https://vast.ai/pricing")
	if err != nil {
		return -1, err
	}

	priceHourFloat, err := strconv.ParseFloat(priceHour, 64)
	if err != nil {
		return -1, err
	}

	return priceHourFloat, nil
}

func vastPriceHandler(e *colly.HTMLElement, priceHour *string) {
	rowText := e.Text
	if strings.Contains(strings.ToLower(rowText), "h200") {
		e.ForEach("td", func(i int, td *colly.HTMLElement) {
			if strings.HasPrefix(td.Text, "$") && *priceHour == "" {
				*priceHour = strings.Split(td.Text, "$")[1]
			}
		})
	}
}