package routes

import (
	"net/http"
	"targon-oracle/internal/scraping"

	"github.com/labstack/echo/v4"
)


func AveragePriceHandler(c echo.Context) error {
	togetherPrice := scraping.TogetherScraper()
	vastPrice := scraping.VastScraper()
	crusoePrice := scraping.CrusoeScraper()
	shadeFormPrice := scraping.ShadeFormScraper()
	coreweavePrice := scraping.CoreweaveScraper()
	nebiusPrice := scraping.NebiusScraper()

	avg := (togetherPrice + vastPrice + crusoePrice + shadeFormPrice + coreweavePrice + nebiusPrice) / 6

	return c.JSON(http.StatusOK, map[string]interface{}{
		"gpu_type":      "h200",
		"average_price": avg,
	})
}
