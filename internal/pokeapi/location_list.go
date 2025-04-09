package pokeapi

import (
	"encoding/json"
	"io"
	"fmt"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var body []byte
	body, available := c.cache.Get(url)

	if !available {
		// fmt.Println("No Cache")
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, fmt.Errorf("request error: %w", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, fmt.Errorf("Get error: %w", err)
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return RespShallowLocations{}, fmt.Errorf("Body error: %w", err)
		}
		if res.StatusCode > 299 {
			return RespShallowLocations{}, fmt.Errorf("Server error: %v - %s", res.StatusCode, res.Status)
		}

		c.cache.Add(url, body)
	}

	jsonData := RespShallowLocations{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return RespShallowLocations{}, fmt.Errorf("JSON error: %w", err)
	}

	return jsonData, nil
}