package server

import (
	"net/http"
	"parking-lot/pkg/util"
)

func WriteFailResponse(w http.ResponseWriter, status int, err error, errorMsg string) {
	util.ErrorHandler(err, errorMsg)
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(status)
	errorMsg = errorMsg + "\n\n"
	w.Write([]byte(errorMsg))
}

func WriteSuccessResponse(w http.ResponseWriter, data string) {
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	data = data + "\n\n"
	w.Write([]byte(data))
}
