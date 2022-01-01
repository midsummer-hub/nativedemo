package main

import (
	"fmt"
	"github.com/nativedemo/httpdemo/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/callInfo", handler.AccessLogging(handler.CallInfoHandler))
	http.HandleFunc("/healthz", handler.AccessLogging(handler.HalthzHandler))

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("Faild to start handler,err:%s", err.Error())
	}
}
