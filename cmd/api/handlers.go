package main

import (
	"encoding/json"
	"net/http"
)



func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {

	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	- = app.writeJSON(w, httpStatusOK, paypayload){
		
	}
}
