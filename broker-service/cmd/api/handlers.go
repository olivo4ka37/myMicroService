package main

import (
	"fmt"
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	err := app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		fmt.Println(fmt.Errorf("Error while read json in broker-service/cmd/api/handlers.go ", err))
	}
}
