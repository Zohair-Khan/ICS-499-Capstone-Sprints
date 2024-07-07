package main

import (
	"crypto/tls"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type application struct {
	templateCache map[string]*template.Template
	log           *slog.Logger
}

func main() {

	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	tc, err := newTemplateCache()

	if err != nil {
		log.Warn(err.Error())
		os.Exit(1)
	}

	app := &application{
		templateCache: tc,
		log:           log,
	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         ":443",
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(log.Handler(), slog.LevelInfo),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Info("starting server", "addr", srv.Addr)
	if err := srv.ListenAndServeTLS("/etc/letsencrypt/live/ph-notes.com/cert.pem", "/etc/letsencrypt/live/ph-notes.com/privkey.pem"); err != nil {
		app.log.Warn(err.Error())
		os.Exit(1)
	}
}
