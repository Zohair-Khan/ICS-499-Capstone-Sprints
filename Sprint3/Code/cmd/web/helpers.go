package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const usernameContextKey = contextKey("username")

func getDSN() string {

	dbname := os.Getenv("DB")
	dbuser := os.Getenv("DBUSER")
	dbpassword := os.Getenv("DBPASSWORD")

	if dbuser == "" || dbname == "" || dbpassword == "" {
		return "phadmin:teambadass@/phtestserver?parseTime=true"
	}
	var sb strings.Builder

	sb.WriteString(dbuser)
	sb.WriteString(":")
	sb.WriteString(dbpassword)
	sb.WriteString("@/")
	sb.WriteString(dbname)
	sb.WriteString("?parseTime=true")

	return sb.String()
}
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	method := r.Method
	uri := r.URL.RequestURI()

	app.log.Error(err.Error(), "method", method, "uri", uri)

	app.clientError(w, http.StatusInternalServerError)

}
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, pageName string, data data) {

	// Access the page from the cache
	templateSet, ok := app.templateCache[pageName]
	// If it doesn't exist, warn
	if !ok {
		err := fmt.Errorf("the template %s does not exist", pageName)
		app.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	err := templateSet.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)

	buf.WriteTo(w)

}

func (app *application) HTMLTimeToGoTime(inputDate string, inputTime string) (time.Time, error) {

	dates := strings.Split(inputDate, "-")

	year, err := strconv.Atoi(dates[0])

	if err != nil {
		return time.Time{}, err
	}

	month, err := strconv.Atoi(dates[1])

	if err != nil {

		return time.Time{}, err
	}

	day, err := strconv.Atoi(dates[2])

	if err != nil {

		return time.Time{}, err
	}

	times := strings.Split(inputTime, ":")

	hour, err := strconv.Atoi(times[0])

	if err != nil {

		return time.Time{}, err
	}

	minute, err := strconv.Atoi(times[1])

	if err != nil {

		return time.Time{}, err
	}

	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC), nil
}

func (app *application) validateProvider(r *http.Request) (*http.Request, error) {

	c, err := r.Cookie("token")
	if err != nil {
		return r, err
	}

	tkn, err := jwt.ParseWithClaims(c.Value, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return jwtKey, nil
	})

	if err != nil {
		return r, err
	}

	if !tkn.Valid {
		return r, jwt.ErrInvalidKey
	}

	username, err := tkn.Claims.GetSubject()

	if err != nil {
		return r, jwt.ErrTokenInvalidSubject
	}

	app.log.Info(fmt.Sprintf("Username after helper function: %s", username))
	return r.Clone(context.WithValue(r.Context(), usernameContextKey, username)), nil

}

func (app *application) getTemplateData(r *http.Request) data {

	_, ok := r.Context().Value(usernameContextKey).(string)

	if !ok {
		return data{IsLoggedIn: false}
	}
	return data{IsLoggedIn: true}
}
