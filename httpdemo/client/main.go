package main

import (
	"fmt"
	"net/http"
)

func main() {
	addr := "http://localhost:9090/callInfo"
	client := &http.Client{}

	request,_ := http.NewRequest("GET", addr, nil)
	request.Header.Set("X-Token", "123456")
	request.Header.Set("Cookie", "xxx")
	request.Header.Set("User-Agent", "anonymity")
	response,_ := client.Do(request)
	defer response.Body.Close()
	fmt.Printf("Response Info\n" +
		"Head[VERSION]:%s\n" +
		"Head[X-Token]:%s\n" +
		"Head[Cookie]:%s\n" +
		"Head[user-agent]:%s\n" +
		"StatusCode:%d\n", response.Header["VERSION"],response.Header["X-Token"],response.Header["Cookie"],response.Header["User-Agent"],response.StatusCode)
}