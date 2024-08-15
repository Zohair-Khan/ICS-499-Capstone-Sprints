package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func (app *application) home() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		data := app.getTemplateData(r)
		app.render(w, r, http.StatusOK, "home.html", data)

	})
}

func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
	data := app.getTemplateData(r)

	if data.AuthLevel > 0 {
		http.Redirect(w, r, "/notes/view", http.StatusSeeOther)
	}
	app.render(w, r, http.StatusOK, "login.html", data)
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")

	if errors.Is(err, http.ErrNoCookie) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	c.Expires = time.Now()

	// Override token cookie
	http.SetCookie(w, c)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *application) getAdminNotesView(w http.ResponseWriter, r *http.Request) {
	data := app.getTemplateData(r)
	app.render(w, r, http.StatusOK, "notes-admin.html", data)
}

func (app *application) getAdminNoteView(w http.ResponseWriter, r *http.Request) {
	data := app.getTemplateData(r)
	app.render(w, r, http.StatusOK, "note-admin.html", data)
}

func (app *application) getNotesView(w http.ResponseWriter, r *http.Request) {
	/*
		username, ok := r.Context().Value(usernameContextKey).(string)

		if !ok {
			app.serverError(w, r, fmt.Errorf("no username stored in context"))
			return
		}
	*/
	notes, err := app.notes.GetNotes()

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.getTemplateData(r)

	data.Notes = notes

	app.render(w, r, http.StatusOK, "notes.html", data)
}

func (app *application) getNoteView(w http.ResponseWriter, r *http.Request) {
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
			app.clientError(w, r, http.StatusBadRequest)
			return
		}
		app.serverError(w, r, err)
		return
	}

	data := app.getTemplateData(r)

	data.Note = note
	/* TODO: reinstate code, but it's not recognizing the user with the note
	username, ok := r.Context().Value(usernameContextKey).(string)

	if !ok {
		app.serverError(w, r, fmt.Errorf("couldn't parse username from context"))
		return
	}


	providerID, err := app.users.GetID(username)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

		if data.Note.ProviderID != providerID {
			app.clientError(w, r, http.StatusForbidden)
			return
		}
	*/
	app.render(w, r, http.StatusOK, "note.html", data)
}

type NoteCreateForm struct {
	Patient     string            `schema:"patient"`
	Patients    []string          `schema:"-"`
	Service     string            `schema:"service"`
	Services    map[string]string `schema:"-"`
	ServiceDate string            `schema:"serviceDate"`
	StartTime   string            `schema:"startTime"`
	EndTime     string            `schema:"endTime"`
	Summary     string            `schema:"summary"`
}

func (app *application) postNoteCreate(w http.ResponseWriter, r *http.Request) {
	// Parse form
	err := r.ParseForm()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	form := NoteCreateForm{}
	// Decode post form into struct
	err = decoder.Decode(&form, r.PostForm)

	if err != nil {
		app.serverError(w, r, err)
		return
	}
	errors := map[string]string{}
	if form.Patient == "none" {
		errors["Patient"] = "Select a patient"
	}

	if form.Service == "none" {
		errors["Service"] = "Select a service"
	}
	// DATE
	if form.ServiceDate == "" {
		errors["ServiceDate"] = "Enter a date of service"
	}

	date, err := time.Parse("2006-01-02", form.ServiceDate)

	if err != nil {
		errors["ServiceDate"] = "Enter a date in the form YYYY-MM-DD"
	}

	if date.After(time.Now()) {
		errors["ServiceDate"] = "Select a date not in the future"
	}

	// TIME
	if form.StartTime == "" {
		errors["StartTime"] = "Enter a valid start time"
	}
	if form.EndTime == "" {
		errors["EndTime"] = "Enter a valid end time"
	}

	end, err := time.Parse("15:04", form.EndTime)
	if err != nil {
		errors["EndTime"] = "Please enter a time in the form HH:MM"
	}

	start, err := time.Parse("15:04", form.StartTime)

	if err != nil {
		errors["StartTime"] = "Please enter a time in the form HH:MM"
	}

	if end.Before(start) {
		errors["EndTime"] = "End time cannot be after start time"
	}

	if form.Summary == "" {
		errors["Summary"] = "Summary cannot be blank"
	}
	// Get patientID from initials
	// NOTE: doing it this way to prevent storage of patient id

	// Get service type

	if len(errors) != 0 {
		fmt.Printf("Number of errors: %d\n", len(errors))
		patients, err := app.patients.GetAll()
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		form.Patients = make([]string, 0)

		// Store their initials as a string and add to form data
		for _, patient := range patients {
			b := strings.Builder{}
			fmt.Fprintf(&b, "%s %s", patient.FirstInitials, patient.LastInitials)
			form.Patients = append(form.Patients, b.String())
		}
		form.Services = map[string]string{
			"general":    "General",
			"individual": "Individual",
			"family":     "Family",
			"group":      "Group",
		}
		data := app.getTemplateData(r)
		data.Form = form
		data.Errors = errors
		app.render(w, r, http.StatusOK, "add-note.html", data)
		return
	}

	var patientID int
	if form.Patient != "" {

		patientInitials := strings.Split(form.Patient, " ")

		patientFirst := patientInitials[0]

		patientLast := patientInitials[1]

		patientID, err = app.patients.GetID(patientFirst, patientLast)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
	}
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

	newID, err := app.notes.Insert(providerID, patientID, form.Service, form.ServiceDate, form.StartTime, form.EndTime, form.Summary)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/note/view/%d", newID), http.StatusSeeOther)

	// upon successful entry, redirect to new note view

}

func (app *application) getNoteCreate(w http.ResponseWriter, r *http.Request) {
	data := app.getTemplateData(r)

	form := NoteCreateForm{}

	// Get list of patients
	patients, err := app.patients.GetAll()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	form.Patients = make([]string, 0)

	// Store their initials as a string and add to form data
	for _, patient := range patients {
		b := strings.Builder{}
		fmt.Fprintf(&b, "%s %s", patient.FirstInitials, patient.LastInitials)
		form.Patients = append(form.Patients, b.String())
	}
	form.Patient = "none"
	form.Services = map[string]string{
		"general":    "General",
		"individual": "Individual",
		"family":     "Family",
		"group":      "Group",
	}
	form.Service = "none"

	form.ServiceDate = time.Now().Format("2006-01-02")

	form.StartTime = time.Now().Format("15:04")
	form.EndTime = time.Now().Format("15:04")

	data.Form = form

	app.render(w, r, http.StatusOK, "add-note.html", data)
}

func (app *application) postNoteEdit(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		app.clientError(w, r, http.StatusBadRequest)
		return
	}

	err = r.ParseForm()

	if err != nil {
		app.clientError(w, r, http.StatusBadRequest)
		return
	}

	status := r.PostForm.Get("status")

	if status == "" {
		app.clientError(w, r, http.StatusBadRequest)
		return
	}

	err = app.notes.UpdateStatus(id, status)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/note/view/%d", id), http.StatusSeeOther)

}
