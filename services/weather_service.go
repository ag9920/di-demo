package services

import (
	"io/ioutil"
	"net/http"
)

type WeatherService struct {
}

func (ws *WeatherService) Weather(city string) (*string, error) {
	response, err := http.Get("https://wttr.in/" + city)
	if err != nil {
		return nil, err
	}
	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	weather := string(all)
	return &weather, nil
}
