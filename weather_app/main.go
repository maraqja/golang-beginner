package main

import (
	"flag"
	"fmt"
	"weather_app/geo"
	"weather_app/weather"
)

func main() {
	city := flag.String("city", "", "Город") // хотим забирать из строки флаги (не работает до вызова parse) - go run . --name = "Kaka"
	format := flag.Int("format", 1, "Формат вывода погоды")
	flag.Parse() // после описания всех флагов вызываем

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(*geoData)
	weatherData, err := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)

}
