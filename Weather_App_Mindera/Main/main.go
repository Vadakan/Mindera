package main

import (
	"fmt"
	"github.com/Weather_App_Mindera"
	"sync"
)
var wg sync.WaitGroup
var Weatherstack,OpenWeather string

func main(){

    fmt.Println("Main started...")
    wg.Add(2)
    go func(){
		Weatherstack = Weather_App_Mindera.GetWeatherDetails()
		wg.Done()
    }()
	go func(){
	   OpenWeather  = Weather_App_Mindera.GetOpenWeather()
	wg.Done()
	}()

    wg.Wait()
	if Weatherstack != " "{
	    fmt.Println(Weatherstack)
	}else{
	    fmt.Println(OpenWeather)
	}

}
