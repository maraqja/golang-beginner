package geo_test

import (
	"testing"
	"weather_app/geo"
)

func TestGetMyLocation(t *testing.T) { // t - позволяет управлять контекстом теста
	// Arange - подготовка теста: expected результат, доп подготовкаm данные для функции
	city := "Moscow"
	expected := geo.GeoData{
		City: "London",
	}
	// Act - выполнение функции
	got, err := geo.GetMyLocation(city)
	// Assert - проверка результата
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Expected %s, got %s", expected.City, got.City)
	}
}

// go test geo/geo_test.go
