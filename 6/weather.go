package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Main struct {
	Temp float64 `json:"temp"`
}

type Response struct {
	Main Main `json:"main"`
}

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type RealHTTPClient struct {
	Client *http.Client
}

func (c *RealHTTPClient) Get(url string) (*http.Response, error) {
	return c.Client.Get(url)
}

func GetWeather(client HTTPClient, city string, apiKey string) (float64, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	var weatherResp Response
	err = json.NewDecoder(resp.Body).Decode(&weatherResp)
	if err != nil {
		return 0, err
	}

	return weatherResp.Main.Temp, nil
}
