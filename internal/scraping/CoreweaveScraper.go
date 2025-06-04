package scraping

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func CoreweaveScraper(collyAgent *colly.Collector) (float64, error) {
	c := collyAgent

	var priceHour string

	c.OnHTML("div.table-row.w-dyn-item.kubernetes-gpu-pricing", func(e *colly.HTMLElement) {
		coreweavePriceHandler(e, &priceHour)
	})

	err := c.Visit("https://www.coreweave.com/pricing")
	if err != nil {
		return -1, err
	}

	priceHourFloat, err := strconv.ParseFloat(priceHour, 64)
	if err != nil {
		return -1, err
	}

	return priceHourFloat / 8, nil
}

func coreweavePriceHandler(e *colly.HTMLElement, priceHour *string) {
	text := e.Text
	if strings.Contains(text, "HGX H200") {
		re := regexp.MustCompile(`\$(\d+\.\d+)`)
		match := re.FindStringSubmatch(text)
		if len(match) > 1 {
			*priceHour = match[1]
		}
	}
}
