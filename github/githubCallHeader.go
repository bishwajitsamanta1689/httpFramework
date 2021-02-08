package main

import (
	"log"
	"net/http"
)

func main() {
	httpMethod := "GET"
	url := "https://api.github.com"
	client := http.Client{}
	request, err := http.NewRequest(httpMethod, url, nil)
	request.Header.Set("accept", "application/json")
	mustt(err)
	resp, err := client.Do(request)
	defer resp.Body.Close()
	mustt(err)

}

func mustt(err error) {
	if err != nil {
		log.Println(err)
	}
}
