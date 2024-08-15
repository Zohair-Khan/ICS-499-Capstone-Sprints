package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	// Serving style and scripts as files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Public
	mux.Handle("/", app.home()) // Is this easier to read?

	mux.HandleFunc("/logout", app.logout)

	mux.HandleFunc("POST /login", app.postLogin)

	mux.HandleFunc("GET /login", app.getLogin)

	mux.Handle("POST /note/edit/{id}", app.registerAuthorization(app.requireAdmin(http.HandlerFunc(app.postNoteEdit))))
	// Admin
	mux.Handle("/admin/notes/view", app.registerAuthorization(app.requireAdmin(http.HandlerFunc(app.getAdminNotesView))))

	mux.HandleFunc("/admin/note/view/{id}", app.getAdminNoteView)

	// Provider
	mux.Handle("/note/view/{id}", app.registerAuthorization(app.requireProvider(http.HandlerFunc(app.getNoteView))))

	mux.Handle("/notes/view", app.registerAuthorization(app.requireProvider(http.HandlerFunc(app.getNotesView))))

	mux.Handle("POST /note/create", app.registerAuthorization(app.requireProvider(http.HandlerFunc(app.postNoteCreate))))

	mux.Handle("GET /note/create", app.registerAuthorization(app.requireProvider(http.HandlerFunc(app.getNoteCreate))))

	// We're using a closure over commonHeaders.
	// There are certain things we want to respond with not matter the request, so this
	// provides a way to call the commonHeaders function, which in turn calls the appropriate
	// function from mux

	return app.recoverPanic(app.logRequest(commonHeaders(mux)))
}
