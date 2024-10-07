package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weather_app/geo"
)

var ErrorWrongFormat = errors.New("ERROR_WRONG_FORMAT")
var ErrorUrlParse = errors.New("ERROR_URL_PARSE")
var ErrorHttp = errors.New("ERROR_HTTP")
var ErrorNot200 = errors.New("ERROR_NOT_200")
var ErrorReadBody = errors.New("ERROR_READ_BODY")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrorWrongFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City) // преобразует строку в полноценный  url
	if err != nil {
		fmt.Println(err.Error())
		return "", ErrorUrlParse
	}
	params := url.Values{}                   // для объявления query
	params.Add("format", fmt.Sprint(format)) // преобразуем в string
	baseUrl.RawQuery = params.Encode()

	response, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", ErrorHttp
	}
	defer response.Body.Close() // не забывать закрыть чтение body!!! (иначе будет утечка памяти) - с помощью defer откладываем выполнение на после выхода из функции
	if response.StatusCode != 200 {
		fmt.Println(err.Error())
		return "", ErrorNot200
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", ErrorReadBody
	}
	return string(body), nil
}
