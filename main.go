package main

import (
	"fmt"
	"io"
	"net/http"
)


func main() {

	url := "https://api.open-meteo.com/v1/forecast?latitude=35.6895&longitude=139.6917&daily=weather_code&timezone=Asia%2FTokyo"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
