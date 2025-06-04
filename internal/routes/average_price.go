package routes

import (
	"math"
	"net/http"
	"targon-oracle/internal/scraping"

	"sync"

	"github.com/gocolly/colly/v2"
	"github.com/labstack/echo/v4"
)

type scrapeResult struct {
	price float64
	err   error
	name  string
}

func AveragePriceHandler(c echo.Context) error {
	collyAgent := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"))

	scrapers := []struct {
		name string
		fn   func() (float64, error)
	}{
		{"TogetherScraper", func() (float64, error) { return scraping.TogetherScraper(collyAgent) }},
		{"VastScraper", func() (float64, error) { return scraping.VastScraper(collyAgent) }},
		{"CrusoeScraper", func() (float64, error) { return scraping.CrusoeScraper(collyAgent) }},
		{"ShadeFormScraper", scraping.ShadeFormScraper},
		{"CoreweaveScraper", func() (float64, error) { return scraping.CoreweaveScraper(collyAgent) }},
		{"NebiusScraper", func() (float64, error) { return scraping.NebiusScraper(collyAgent) }},
	}

	var wg sync.WaitGroup
	results := make(chan scrapeResult, len(scrapers))

	for _, s := range scrapers {
		wg.Add(1)
		go func(sname string, fn func() (float64, error)) {
			defer wg.Done()
			price, err := fn()
			results <- scrapeResult{price: price, err: err, name: sname}
		}(s.name, s.fn)
	}

	wg.Wait()
	close(results)

	prices := []float64{}
	errors := []string{}

	for r := range results {
		if r.err != nil {
			errors = append(errors, r.name+": "+r.err.Error())
		} else {
			prices = append(prices, r.price)
		}
	}

	var avg float64
	if len(prices) > 0 {
		var sum float64
		for _, p := range prices {
			sum += p
		}
		avg = sum / float64(len(prices))
		avg = math.Round(avg*100) / 100
	} else {
		avg = -1
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"gpu_type":      "h200",
		"average_price": avg,
		"errors":        errors,
	})
}
