package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (app *application) mainHandler(w http.ResponseWriter, r *http.Request) {
	r, _ = app.validateProvider(r)

	data := app.getTemplateData(r)
	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
	r, _ = app.validateProvider(r)
	data := app.getTemplateData(r)
	app.render(w, r, http.StatusOK, "login.html", data)
}
func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")

	if errors.Is(err, http.ErrNoCookie) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	c.Expires = time.Now()

	// Override token cookie
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) progressnoteHandler(w http.ResponseWriter, r *http.Request) {
	data := app.getTemplateData(r)
	app.render(w, r, http.StatusOK, "progressnote.html", data)
}

func (app *application) notesAdmin(w http.ResponseWriter, r *http.Request) {
	data := app.getTemplateData(r)
	app.render(w, r, http.StatusOK, "notes-admin.html", data)
}
func (app *application) noteAdmin(w http.ResponseWriter, r *http.Request) {
	data := app.getTemplateData(r)
	app.render(w, r, http.StatusOK, "note-admin.html", data)
}

func (app *application) notesHandler(w http.ResponseWriter, r *http.Request) {

	username, ok := r.Context().Value(usernameContextKey).(string)

	if !ok {
		app.serverError(w, r, fmt.Errorf("no username stored in context"))
		return
	}
	notes, err := app.notes.GetNotesByProvider(username)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.getTemplateData(r)

	data.Notes = notes

	app.render(w, r, http.StatusOK, "notes.html", data)
}

func (app *application) noteHandler(w http.ResponseWriter, r *http.Request) {
	// We're assuming we've already checked if they have appropriate authorization
	// For now we're passing a dummy variable

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	note, err := app.notes.Get(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.clientError(w, http.StatusBadRequest)
			return
		}
		app.serverError(w, r, err)
		return
	}

	data := app.getTemplateData(r)

	data.Note = note
	username, ok := r.Context().Value(usernameContextKey).(string)

	app.log.Info("Username from context in noteHandler", "username", username)

	if !ok {
		app.serverError(w, r, fmt.Errorf("couldn't parse username from context"))
		return
	}

	if data.Note.Username != username {
		app.clientError(w, http.StatusForbidden)
		return
	}

	app.render(w, r, http.StatusOK, "note.html", data)
}

func (app *application) addNotePost(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	patient := r.PostForm.Get("patient")
	if strings.TrimSpace(patient) == "" {
		app.serverError(w, r, fmt.Errorf("patient field not submitted"))
	}

	service := r.PostForm.Get("service")
	if strings.TrimSpace(service) == "" {
		app.serverError(w, r, fmt.Errorf("service field not submitted"))
	}

	serviceDate := r.PostForm.Get("serviceDate")
	if strings.TrimSpace(serviceDate) == "" {
		app.serverError(w, r, fmt.Errorf("serviceDate field not submitted"))
	}

	startTime := r.PostForm.Get("startTime")
	if strings.TrimSpace(startTime) == "" {
		app.serverError(w, r, fmt.Errorf("startTime field not submitted"))
	}

	endTime := r.PostForm.Get("endTime")
	if strings.TrimSpace(endTime) == "" {
		app.serverError(w, r, fmt.Errorf("endTime field not submitted"))
	}

	summary := r.PostForm.Get("summary")
	if strings.TrimSpace(summary) == "" {
		app.serverError(w, r, fmt.Errorf("summary field not submitted"))
	}

	progress := r.PostForm.Get("progress")
	if strings.TrimSpace(progress) == "" {
		app.serverError(w, r, fmt.Errorf("progress field not submitted"))
	}

	response := r.PostForm.Get("response")
	if strings.TrimSpace(response) == "" {
		app.serverError(w, r, fmt.Errorf("response field not submitted"))
	}

	assessmentStatus := r.PostForm.Get("assessmentStatus")
	if strings.TrimSpace(assessmentStatus) == "" {
		app.serverError(w, r, fmt.Errorf("assessmentStatus field not submitted"))
	}

	riskFactors := r.PostForm.Get("riskFactors")
	if strings.TrimSpace(riskFactors) == "" {
		app.serverError(w, r, fmt.Errorf("riskFactors field not submitted"))
	}

	emergencyInterventions := r.PostForm.Get("emergencyInterventions")
	if strings.TrimSpace(emergencyInterventions) == "" {
		app.serverError(w, r, fmt.Errorf("emergencyInterventions field not submitted"))
	}

	// Get username from context
	provider, ok := r.Context().Value(usernameContextKey).(string)

	if !ok {
		app.serverError(w, r, fmt.Errorf("failed to parse provider from context"))
		return
	}

	providerID, err := app.users.GetID(provider)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	newID, err := app.notes.Insert(providerID, patient, service, serviceDate, startTime, endTime, summary, progress, response, assessmentStatus, riskFactors, emergencyInterventions)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/note/%d", newID), http.StatusSeeOther)

	// upon successful entry, redirect to new note view

}

func (app *application) addNoteGet(w http.ResponseWriter, r *http.Request) {
	data := app.getTemplateData(r)

	username := r.Context().Value("username")

	if username == nil {
		app.log.Error("No context with username provided.")
	}

	app.log.Info(fmt.Sprintf("username: %s", username))

	app.render(w, r, http.StatusOK, "add-note.html", data)
}

func (app *application) addNote1(w http.ResponseWriter, r *http.Request) {

	data := app.getTemplateData(r)

	app.render(w, r, http.StatusOK, "add-note-1.html", data)
}
