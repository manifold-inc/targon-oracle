package scraping

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func CoreweaveScraper() float64 {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	var priceHour string

	c.OnHTML("div.table-row.w-dyn-item.kubernetes-gpu-pricing", func(e *colly.HTMLElement) {
		text := e.Text
		if strings.Contains(text, "HGX H200") {
			re := regexp.MustCompile(`\$(\d+\.\d+)`)
			match := re.FindStringSubmatch(text)
			if len(match) > 1 {
				priceHour = match[1]
			}
		}
	})

	err := c.Visit("https://www.coreweave.com/pricing")
	if err != nil {
		log.Fatal(err)
	}

	priceHourFloat, err := strconv.ParseFloat(priceHour, 64)
	if err != nil {
		log.Fatal(err)
	}

	return priceHourFloat / 8
}
