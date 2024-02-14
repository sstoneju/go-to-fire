package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	method := "POST"
	url := "https://openapivts.koreainvestment.com:29443/oauth2/tokenP"

	payload := strings.NewReader(`{"grant_type": "client_credentials", "appkey": "", "appsecret": ""}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
