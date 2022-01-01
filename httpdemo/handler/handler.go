package handler

import (
	"fmt"
	"net/http"
	"os"
)

func CallInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	w.Header().Set("X-Token", r.Header.Get("X-Token"))
	w.Header().Set("X-Token", r.Header.Get("X-Token"))
	w.Header().Set("Cookie", r.Header.Get("Cookie"))
	w.Header().Set("User-Agent", r.Header.Get("User-Agent"))
	w.WriteHeader(http.StatusOK)
}

func HalthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, http.StatusOK)
}
