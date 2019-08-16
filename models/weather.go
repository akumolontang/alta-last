package models

import (
	"net/http"
	"github.com/labstack/echo"
	"io/ioutil"
	"encoding/json"
	"math"
)


type WeatherCity struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		Pressure  float64 `json:"pressure"`
		Humidity  int     `json:"humidity"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		SeaLevel  float64 `json:"sea_level"`
		GrndLevel float64 `json:"grnd_level"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

// get single weather by city
func GetWeatherController(c echo.Context) error {
		key := "97a0db62b9ec2607eaf18e0dc038ea42"
		city := c.Param("city")
		response, _ := http.Get("https://api.openweathermap.org/data/2.5/weather?appid="+key+"&q="+city)
		responseData, _ := ioutil.ReadAll(response.Body)
	  defer response.Body.Close()
	
	  var TempCity WeatherCity
	  json.Unmarshal(responseData, &TempCity)
  
	  TemperaturCelcius := (math.Round((TempCity.Main.Temp - 273.00)*100)/100)
  
		return c.JSON(http.StatusOK, map[string]interface{}{
	  "city": city,
	  "temperatur": TemperaturCelcius,
	})
}