package server

import (
	"net/http"
	"parking-lot/pkg/util"
)

func WriteFailResponse(w http.ResponseWriter, status int, err error, errorMsg string) {
	util.ErrorHandler(err, errorMsg)
	w.Header().Add("Content-Type", "text/plain")
	errorMsg = errorMsg + "\n"
	w.WriteHeader(status)
	w.Write([]byte(errorMsg))
}

func WriteSuccessResponse(w http.ResponseWriter, data string) {
	w.Header().Add("Content-Type", "text/plain")
	data = data + "\n"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

func WriteBufferResponse(w http.ResponseWriter, status int, err error, response string) {
	if err != nil {
		util.ErrorHandler(err, response)
	}
	w.Header().Add("Content-Type", "text/plain")
	response = response + "\n"
	w.Write([]byte(response))
}
