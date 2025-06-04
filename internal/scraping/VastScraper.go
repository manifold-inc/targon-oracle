package scraping

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func VastScraper() float64 {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	var priceHour string

	c.OnHTML("tbody.css-yp9msd.elmp00v6 tr", func(e *colly.HTMLElement) {
		rowText := e.Text
		if strings.Contains(strings.ToLower(rowText), "h200") {
			e.ForEach("td", func(i int, td *colly.HTMLElement) {
				if strings.HasPrefix(td.Text, "$") && priceHour == "" {
					priceHour = strings.Split(td.Text, "$")[1]
				}
			})
		}
	})

	err := c.Visit("https://vast.ai/pricing")
	if err != nil {
		log.Fatal(err)
	}

	priceHourFloat, err := strconv.ParseFloat(priceHour, 64)
	if err != nil {
		log.Fatal(err)
	}

	return priceHourFloat
}