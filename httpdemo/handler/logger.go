package handler

import (
	"bytes"
	"github.com/nativedemo/httpdemo/utils"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ResponseWithResponse struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func AccessLogging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		logEntry := utils.ALog.WithFields(logrus.Fields{
			"IP":          r.RemoteAddr,
			"Method":      r.Method,
			"Path":        r.RequestURI,
			"Params":       r.URL.Query(),
			"RequestBody": buf.String(),
		})
		rwr := &ResponseWithResponse{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           bytes.Buffer{},
		}
		f.ServeHTTP(rwr, r)
		defer logEntry.WithFields(logrus.Fields{
			"Status":       rwr.statusCode,
			"ResponseBody": rwr.body.String(),
		}).Info()
	}
}

func (r *ResponseWithResponse) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.statusCode = statusCode
}

func (r *ResponseWithResponse) Writer(d []byte) (n int, err error) {
	n, err = r.ResponseWriter.Write(d)
	if err != nil {
		return
	}
	r.body.Write(d)
	return
}
