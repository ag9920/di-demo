package main

import (
	"net/http"

	"github.com/ag9920/di-demo/controllers"

	"github.com/goioc/di"
)

func main() {
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		di.GetInstance("weatherController").(*controllers.WeatherController).Weather(w, r)
	})
	_ = http.ListenAndServe(":8080", nil)
}
