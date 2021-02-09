package main

import (
	"fmt"
	"github.com/httpFramework/goHttp"
)

func main() {
	client := goHttp.New()
	resp, _ := client.Get("https://api.github.com", nil)
	fmt.Println("Status Code:", resp.StatusCode)
}
