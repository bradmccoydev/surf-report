package surf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bradmccoydev/surf-report/model"
)

func GetTides(apiToken string) model.Tide {
	url := "https://api.stormglass.io/v2/tide/sea-level/point?lat=-33.7979&lng=151.2882"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	b := []byte(body)
	var tides model.Tide

	if err := json.Unmarshal(b, &tides); err != nil {
		fmt.Println(err)
	}

	return tides
}

func GetWeather(apiToken string) model.Weather {
	url := "https://api.stormglass.io/v2/weather/point?lat=-33.7979&lng=151.2882&params=airTemperature,cloudCover,swellDirection,swellHeight,swellPeriod,waterTemperature,waveDirection,waveHeight"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	b := []byte(body)

	var weather model.Weather

	if err := json.Unmarshal(b, &weather); err != nil {
		fmt.Println(err)
	}

	return weather
}
