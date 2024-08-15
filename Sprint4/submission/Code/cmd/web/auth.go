package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO: use a env variable to get this from the system
var (
	jwtKey = []byte("secret-key")
)

type CustomClaims struct {
	Username  string `json:"username"`
	AuthLevel int    `json:"authlevel"`
	jwt.RegisteredClaims
}

func (c CustomClaims) Validate() error {
	if c.Username == "" {
		return fmt.Errorf("no username in token")
	}
	if c.AuthLevel < 0 {
		return fmt.Errorf("invalid auth level")
	}
	return nil
}

func (app *application) postLogin(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		app.log.Error("Couldn't parse form.")
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	errors := map[string]string{}

	username := r.PostForm.Get("user")
	if strings.Compare(username, "") == 0 {
		errors["Username"] = "Username cannot be blank"
	}
	password := r.PostForm.Get("pass")
	if strings.Compare(password, "") == 0 {
		errors["Password"] = "Password cannot be blank"
	}
	// Get hashed password from DB and compare with submitted hashed password
	authLevel, err := app.users.Validate(username, password)

	if err != nil {
		errors["General"] = "Either username or password are incorrect"
	}
	// Register the form with the template data for populating errors if any
	data := app.getTemplateData(r)

	data.Errors = errors
	// If errors exist, re-render the login page with them populated
	if len(data.Errors) != 0 {
		app.render(w, r, http.StatusBadRequest, "login.html", data)
		return
	}
	// If we've made it this far, we know the user is legit and can create a cookie
	expirationTime := time.Now().Add(time.Hour)

	claims := &CustomClaims{
		username,
		authLevel,
		jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expirationTime), IssuedAt: jwt.NewNumericDate(time.Now())},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	http.Redirect(w, r, "/notes/view", http.StatusSeeOther)

}
