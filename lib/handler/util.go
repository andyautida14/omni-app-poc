package handler

import (
	"log"
	"log/slog"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func HandleInternalServerError(w http.ResponseWriter, _ *http.Request, err error) {
	slog.Error("an unexpected error occured:" + err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func HandleNotFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func DSMust[T any](ds interface{}, err error) T {
	if err != nil {
		log.Fatalf("Datastore retrieval error: " + err.Error())
	}

	dsT, ok := ds.(T)
	if !ok {
		log.Fatalf("Datastore type invalid")
	}
	return dsT
}
