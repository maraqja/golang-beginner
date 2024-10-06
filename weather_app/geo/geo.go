package geo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json: "city"`
}

func GetMyLocation(city string) (*GeoData, error) { // если не указан city (""), то получаем с помощью стороннего API город локации вызова
	if city != "" {
		return &GeoData{
			City: city,
		}, nil
	}
	response, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("NOT_200")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}