package scraping

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
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

func CrusoeScraper() float64 {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	var priceHour string

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		rowText := e.Text
		if strings.Contains(strings.ToLower(rowText), "h200") {
			priceHour = strings.Split(rowText, "$")[1]
			priceHour = strings.Fields(priceHour)[0]
		}
	})

	err := c.Visit("https://crusoe.ai/cloud/pricing?instance=3&numOfGpus=0")
	if err != nil {
		log.Fatal(err)
	}

	priceHourFloat, err := strconv.ParseFloat(priceHour, 64)
	if err != nil {
		log.Fatal(err)
	}

	return priceHourFloat
}

func ShadeFormScraper() float64 {
	resp, err := http.Get("https://api.shadeform.ai/v1/instances/types")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data struct {
		InstanceTypes []struct {
			Configuration struct {
				GpuType      string `json:"gpu_type"`
				NumGpus      int    `json:"num_gpus"`
				Interconnect string `json:"interconnect"`
				Region       string `json:"region"`
			} `json:"configuration"`
			HourlyPrice  float64 `json:"hourly_price"`
			Availability []struct {
				Region      string `json:"region"`
				Available   bool   `json:"available"`
				DisplayName string `json:"display_name"`
			} `json:"availability"`
		} `json:"instance_types"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	for _, inst := range data.InstanceTypes {
		if inst.Configuration.GpuType == "H200" && inst.HourlyPrice == 245 {
			allUnavailable := true
			for _, region := range inst.Availability {
				if region.Available {
					allUnavailable = false
					break
				}
			}
			if allUnavailable {
				return inst.HourlyPrice / 100
			}
		}
	}

	return 0
}

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
