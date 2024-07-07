package main

import (
	"net/http"
)

func (app *application) mainHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.html")
}

func (app *application) progressnoteHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "progressnote.html")
}

func (app *application) notesAdmin(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "notes-admin.html")
}
func (app *application) noteAdmin(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "note-admin.html")
}
func (app *application) notes(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "notes.html")
}

func (app *application) note(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "note.html")
}

func (app *application) addNote1(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "add-note-1.html")
}

func (app *application) addNote2(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "add-note-2.html")
}
func (app *application) addNote3(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "add-note-3.html")
}
func (app *application) addNote4(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "add-note-4.html")
}
func (app *application) addNote5(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "add-note-5.html")
}
func (app *application) addNote6(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "add-note-6.html")
}
func (app *application) addNote7(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "add-note-7.html")
}
