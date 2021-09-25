//package will get Weatherstack API URL to get Wind speed and Temperature
package Weather_App_Mindera

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//WSResult response details from Weather Stack API
type WSResult struct {
	Request struct {
		Type     string `json:"type"`
		Query    string `json:"query"`
		Language string `json:"language"`
		Unit     string `json:"unit"`
	} `json:"request"`
	Location struct {
		Name           string `json:"name"`
		Country        string `json:"country"`
		Region         string `json:"region"`
		Lat            string `json:"lat"`
		Lon            string `json:"lon"`
		TimezoneID     string `json:"timezone_id"`
		Localtime      string `json:"localtime"`
		LocaltimeEpoch int    `json:"localtime_epoch"`
		UtcOffset      string `json:"utc_offset"`
	} `json:"location"`
	Current struct {
		ObservationTime     string   `json:"observation_time"`
		Temperature         int      `json:"temperature"`
		WeatherCode         int      `json:"weather_code"`
		WeatherIcons        []string `json:"weather_icons"`
		WeatherDescriptions []string `json:"weather_descriptions"`
		WindSpeed           int      `json:"wind_speed"`
		WindDegree          int      `json:"wind_degree"`
		WindDir             string   `json:"wind_dir"`
		Pressure            int      `json:"pressure"`
		Precip              int      `json:"precip"`
		Humidity            int      `json:"humidity"`
		Cloudcover          int      `json:"cloudcover"`
		Feelslike           int      `json:"feelslike"`
		UvIndex             int      `json:"uv_index"`
		Visibility          int      `json:"visibility"`
	} `json:"current"`
}

//WSres Object for Response details from weather stack API
var WSres WSResult

//WSDisplay the output
type WSDisplay struct {
	WindSpeed          int `json:wind_speed`
	TemperatureDegrees int `json:temperature_degrees`
}

func GetWeatherDetails() string {

	WSfullurl := `http://api.weatherstack.com/current?access_key=&query=Melbourne`

	WSreq, WSerr := http.NewRequest("GET", WSfullurl, nil)

	if WSerr != nil {

		panic(WSerr)
	}

	WSclient := &http.Client{}

	WSresp, WSerr1 := WSclient.Do(WSreq)

	if WSerr1 != nil {

		panic(WSerr1)
	}

	defer WSresp.Body.Close()

	WSbody, WSdataReaderr := ioutil.ReadAll(WSresp.Body)
	if WSdataReaderr != nil {
		panic(WSdataReaderr)
	}

	json.Unmarshal(WSbody, &WSres)

	WSInterim_Output := WSDisplay{
		WindSpeed:          WSres.Current.WindSpeed,
		TemperatureDegrees: WSres.Current.Temperature,
	}

	WSFinal_output, WSmarshal_err := json.Marshal(WSInterim_Output)

	if WSmarshal_err != nil {
		panic(WSmarshal_err)
	}
	return string(WSFinal_output)
}
