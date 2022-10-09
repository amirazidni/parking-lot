package util

import "log"

func ErrorHandler(err error, message string) bool {
	if err != nil {
		log.Default().Printf(message + ". Cause: " + err.Error())
		return true
	}
	return false
}

func ErrorHandlerFatal(err error, message string) bool {
	if err != nil {
		log.Default().Fatal(message + ". Cause: " + err.Error())
		return true
	}
	return false
}
