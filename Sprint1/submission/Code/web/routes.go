package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	//This is the same as mux.Handle("/", http.Handler(app.mainHandler))
	mux.HandleFunc("/", app.mainHandler)
	mux.HandleFunc("/progressnote", app.progressnoteHandler)
	mux.HandleFunc("/notes-admin", app.notesAdmin)
	mux.HandleFunc("/note-admin/{id}", app.noteAdmin)
	mux.HandleFunc("/notes", app.notes)
	mux.HandleFunc("/note/{id}", app.note)

	//Progress Notes Path
	mux.HandleFunc("/add-note-1", app.addNote1)
	mux.HandleFunc("/add-note-2", app.addNote2)
	mux.HandleFunc("/add-note-3", app.addNote3)
	mux.HandleFunc("/add-note-4", app.addNote4)
	mux.HandleFunc("/add-note-5", app.addNote5)
	mux.HandleFunc("/add-note-6", app.addNote6)
	mux.HandleFunc("/add-note-7", app.addNote7)
	return mux
}
