package main

import (
	"reflect"

	"github.com/ag9920/di-demo/controllers"
	"github.com/ag9920/di-demo/services"

	"github.com/goioc/di"
)

func init() {
	_, _ = di.RegisterBean("weatherService", reflect.TypeOf((*services.WeatherService)(nil)))
	_, _ = di.RegisterBean("weatherController", reflect.TypeOf((*controllers.WeatherController)(nil)))
	_ = di.InitializeContainer()
}
