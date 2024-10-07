package weather_test

import (
	"strings"
	"testing"
	"weather_app/geo"
	"weather_app/weather"
)

func TestGetWeather(t *testing.T) {
	expected := "Moscow"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3 // в этом формате выводится город в ответе

	result, err := weather.GetWeather(geoData, format)

	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(result, expected) {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestGetWeatherWrongFormat(t *testing.T) {
	expected := "Moscow"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 111

	_, err := weather.GetWeather(geoData, format)
	if err != weather.ErrorWrongFormat {
		t.Errorf("Expected %s, got %s", weather.ErrorWrongFormat, err)
	}

}
