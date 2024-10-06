package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weather_app/geo"
)

func GetWeather(geo geo.GeoData, format int ) string {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City) // преобразует строку в полноценный  url
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	params := url.Values{} // для объявления query
	params.Add("format", fmt.Sprint(format)) // преобразуем в string
	baseUrl.RawQuery = params.Encode()

	response, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	if response.StatusCode != 200 {
		fmt.Println(err.Error())
		return ""
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(body)
}