package main

import "net/http"

func (app *application) getAppStatus(w http.ResponseWriter, r *http.Request) {
	status := AppStatus{
		Status:      "Available",
		Enviornment: app.config.env,
		Version:     version,
	}
	app.writeResponse(w, http.StatusOK, status)
}
