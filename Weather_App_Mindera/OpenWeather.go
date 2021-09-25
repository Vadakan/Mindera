package Weather_App_Mindera

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Result response details from Weather Stack API
type OpenWeatherReponse struct {
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
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
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
//Object for Response details from weather stack API
var OWres OpenWeatherReponse

//Display the output
type OWDisplay struct{
	Wind_speed int `json:wind_speed`
	Temperature_degrees int  `json:temperature_degrees`
}

func GetOpenWeather() string{
	fullurl := `http://api.openweathermap.org/data/2.5/weather?q=melbourne&appid=`

	OWreq, OWerr := http.NewRequest("GET", fullurl, nil)

	if OWerr != nil {

		panic(OWerr)
	}

	OWclient := &http.Client{}

	OWresp, OWerr1 := OWclient.Do(OWreq)

	if OWerr1 != nil {

		panic(OWerr1)
	}

	defer OWresp.Body.Close()


	OWbody, OWdataReaderr := ioutil.ReadAll(OWresp.Body)
	if OWdataReaderr != nil {
		panic(OWdataReaderr)
	}

	json.Unmarshal(OWbody,&OWres)

	OWInterim_Output := OWDisplay{
		Wind_speed : int(OWres.Main.Temp),
		Temperature_degrees : int(OWres.Wind.Speed),
	}

	OWFinal_output, OWmarshal_err :=  json.Marshal(OWInterim_Output)

	if OWmarshal_err != nil{
		panic(OWmarshal_err)
	}
	return string(OWFinal_output)
}