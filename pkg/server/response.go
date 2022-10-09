package server

import (
	"net/http"
)

func WriteFailResponse(w http.ResponseWriter, status int, errorMsg string) {
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
