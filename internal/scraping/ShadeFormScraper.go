package scraping

import (
	"encoding/json"
	"io"
	"net/http"
)

func ShadeFormScraper() (float64, error) {
	resp, err := http.Get("https://api.shadeform.ai/v1/instances/types")
	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, err
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
		return -1, err
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
				return inst.HourlyPrice / 100, nil
			}
		}
	}

	return -1, nil
}