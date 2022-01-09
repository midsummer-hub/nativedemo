package main

import (
	"fmt"
	"github.com/nativedemo/httpdemo/handler"
	log "github.com/nativedemo/httpdemo/utils"
	"net/http"
)

func main() {
	log.ALog.Println("server started")
	http.HandleFunc("/callInfo", handler.AccessLogging(handler.CallInfoHandler))
	http.HandleFunc("/healthz", handler.AccessLogging(handler.HalthzHandler))

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("Faild to start handler,err:%s", err.Error())
	}
}
