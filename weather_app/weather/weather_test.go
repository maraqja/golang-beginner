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

var testCases = []struct {
	name   string
	format int
}{
	{"Too big format", 11111},
	{"0", 0},
	{"Minus format", -1}, // можно дополнительно для каждого тест кейса объявить expected
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := "Moscow"
			geoData := geo.GeoData{
				City: expected,
			}

			_, err := weather.GetWeather(geoData, tc.format)
			if err != weather.ErrorWrongFormat {
				t.Errorf("Expected %s, got %s", weather.ErrorWrongFormat, err)
			}
		})
	}

}
