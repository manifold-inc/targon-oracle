package scraping

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"regexp"
)

func NebiusScraper() float64 {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	var priceHour string

	c.OnHTML("div.col.col-md-8.col-12.offset-md-1", func(e *colly.HTMLElement) {
		text := e.Text
		if strings.Contains(text, "NVIDIA H200 GPU") {
			re := regexp.MustCompile(`NVIDIA H200 GPU.*?\$(\d+\.\d+)`)
			match := re.FindStringSubmatch(text)
			if len(match) > 1 {
				priceHour = match[1]
			}
		}
	})

	err := c.Visit("https://nebius.com/prices")
	if err != nil {
		log.Fatal(err)
	}

	priceHourFloat, err := strconv.ParseFloat(priceHour, 64)
	if err != nil {
		log.Fatal(err)
	}

	return priceHourFloat
}
