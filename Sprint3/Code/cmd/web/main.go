package main

import (
	"crypto/tls"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/ErichBerger/phtestserver/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	templateCache map[string]*template.Template
	log           *slog.Logger
	users         *models.UserModel
	notes         *models.NoteModel
}

func main() {

	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	tc, err := newTemplateCache()

	if err != nil {
		log.Warn(err.Error())
		os.Exit(1)
	}

	dsn := getDSN()
	println(dsn)
	db, err := models.NewDB(dsn)

	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()
	// if the db is not instantiated, it should exit the program.
	// for right now, until we actually make it, we can just fake the results.

	app := &application{
		templateCache: tc,
		log:           log,
		users:         &models.UserModel{DB: db},
		notes:         &models.NoteModel{DB: db},
	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(log.Handler(), slog.LevelInfo),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Info("starting server", "addr", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		app.log.Error(err.Error())
		os.Exit(1)
	}
}
