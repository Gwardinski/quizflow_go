package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeResponse(w http.ResponseWriter, status int, data interface{}) {
	err := app.writeJSON(w, status, data, "data")
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
	}
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *application) writeError(w http.ResponseWriter, e error, status int) error {
	type jsonError struct {
		Message string `json:"message"`
	}
	eJson := jsonError{
		Message: e.Error(),
	}
	err := app.writeJSON(w, status, eJson, "error")
	if err != nil {
		return err
	}
	return nil
}
