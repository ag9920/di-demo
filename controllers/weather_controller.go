package controllers

import (
	"net/http"

	"github.com/ag9920/di-demo/services"
)

type WeatherController struct {
	// note that injection works even with unexported fields
	weatherService *services.WeatherService `di.inject:"weatherService"`
}

func (wc *WeatherController) Weather(w http.ResponseWriter, r *http.Request) {
	weather, _ := wc.weatherService.Weather(r.URL.Query().Get("city"))
	_, _ = w.Write([]byte(*weather))
}
