package routes

import (
	"net/http"
	"targon-oracle/internal/scraping"

	"github.com/gocolly/colly/v2"
	"github.com/labstack/echo/v4"
)

func AveragePriceHandler(c echo.Context) error {
	collyAgent := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"))

	prices := []float64{}
	errors := []string{}

	togetherPrice, err := scraping.TogetherScraper(collyAgent)
	if err != nil {
		errors = append(errors, "TogetherScraper: "+ err.Error())
	} else {
		prices = append(prices, togetherPrice)
	}

	vastPrice, err := scraping.VastScraper(collyAgent)
	if err != nil {
		errors = append(errors, "VastScraper: "+ err.Error())
	} else {
		prices = append(prices, vastPrice)
	}

	crusoePrice, err := scraping.CrusoeScraper(collyAgent)
	if err != nil {
		errors = append(errors, "CrusoeScraper: "+ err.Error())
	} else {
		prices = append(prices, crusoePrice)
	}

	shadeFormPrice, err := scraping.ShadeFormScraper()
	if err != nil {
		errors = append(errors, "ShadeFormScraper: "+ err.Error())
	} else {
		prices = append(prices, shadeFormPrice)
	}

	coreweavePrice, err := scraping.CoreweaveScraper(collyAgent)
	if err != nil {
		errors = append(errors, "CoreweaveScraper: "+ err.Error())
	} else {
		prices = append(prices, coreweavePrice)
	}

	nebiusPrice, err := scraping.NebiusScraper(collyAgent)
	if err != nil {
		errors = append(errors, "NebiusScraper: "+ err.Error())
	} else {
		prices = append(prices, nebiusPrice)
	}

	var avg float64
	if len(prices) > 0 {
		var sum float64
		for _, p := range prices {
			sum += p
		}
		avg = sum / float64(len(prices))
	} else {
		avg = -1
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"gpu_type":      "h200",
		"average_price": avg,
		"errors":        errors,
	})
}
