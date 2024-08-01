package main

import (
	"html/template"
	"path/filepath"

	"github.com/ErichBerger/phtestserver/internal/models"
)

type data struct {
	Note       models.Note
	Notes      []models.Note
	User       models.User
	Form       any
	IsLoggedIn bool
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		files := []string{
			"./ui/html/base.html",
			"./ui/html/partials/nav.html",
			page,
		}

		templateSet, err := template.ParseFiles(files...)

		if err != nil {
			return nil, err
		}

		cache[name] = templateSet

	}

	return cache, nil
}
