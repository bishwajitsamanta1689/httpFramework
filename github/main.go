package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://api.github.com"
	client := http.Client{}
	resp, err := client.Get(url)
	must(err)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	must(err)
	fmt.Println(resp.StatusCode)
	fmt.Println(string(data))

}
func must(err error) {
	if err != nil {
		log.Println(err)
	}
}
