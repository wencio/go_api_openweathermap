package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Main struct {
	Temp          float64 `json:"temp"`
	TempFeelsLike float64 `json:"feels_like"`
	TempMin       float64 `json:"temp_min"`
	TempMax       float64 `json:"temp_max"`
	Pressure      float64 `json:"pressure"`
}

type WeatherResponse struct {
	Name string `json:"name"`
	Main Main   `json:"main"`
}

func main() {

	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Miami&appid=4f44e31055301a414199df6529ee49f0&units=metric")
	if err != nil {
		log.Fatal(err)
	}

	bytes, errRead := ioutil.ReadAll(response.Body)

	defer func() {
		e := response.Body.Close()
		if e != nil {
			log.Fatal(e)
		}
	}()

	if errRead != nil {
		log.Fatal(errRead)
	}

	log.Print(string(bytes))

	var weatherResponse WeatherResponse

	err = json.Unmarshal(bytes, &weatherResponse)

	if err != nil {

		log.Fatal(err)
	}

	log.Printf("%+v", weatherResponse)
}
