package main

import (
	"github.com/olivo4ka37/myicroervice/logger-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// read json into var
	var requestPayLoad JSONPayload
	_ = app.readJSON(w, r, requestPayLoad)

	// insert the data
	event := data.LogEntry{
		Name: requestPayLoad.Name,
		Data: requestPayLoad.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
	}

	resp := jsonResponse{
		Error:   false,
		Message: "Logged!",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}
