package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	// TODO: handle error
	_ = app.WriteJson(w, http.StatusOK, payload)
}
