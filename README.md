# This app will refer the Weatherstack and openweatherMap API's and fetch the current Wind speed and Temperature in JSON load.

**Project inlcudes:**

1) Go MOD file          - Handling package dependencies
2) OpenWeather.go       - Function To fetch the current Temperature and windspeed of Melbourne city using OpenweatherMAP API
3) OpenWeather_test.go  - Benchmark testing script to test the OpenWeather.go function
4) Weather_app.go       - Function To fetch the current Temperature and windspeed of Melbourne city using Weatherstack API
5) Weather_app_test.go  - Benchmark testing script to test the Weather_app.go function
6) Main.go              - Executable file to get the current wind speed and Temperature of Melbourne City.

Note: 
1) I intentionally removed the API key from app because i have only free subscription so if any accidental run of Benchmark script will lead me to pay for the app.
2) Didnt include Example and Tabletests since we are hardcoding Melbourne city name as asked in requirement document.
