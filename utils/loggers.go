package utils

import (
	"net/http"
)

func RespondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
	return
}

func RespondNotFound(message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(message))
	return

}

func RespondNotContent(message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(message))

}
