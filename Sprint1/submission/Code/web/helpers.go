package main

import (
	"net/http"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, pageName string) {
	templateSet, ok := app.templateCache[pageName]

	if !ok {

		app.log.Warn("Template not found", "name", pageName)
		return
	}

	w.WriteHeader(status)

	err := templateSet.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.log.Warn("template not executed", "name", pageName)
		return
	}

}
