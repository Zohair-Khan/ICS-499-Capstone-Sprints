package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const usernameContextKey = contextKey("username")

const authLevelContextKey = contextKey("authLevel")

const authLevelProvider = 1
const authLevelAdmin = 2

func commonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set some common headers here
		w.Header().Set("Server", "Go")
		// w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")
		next.ServeHTTP(w, r)

		// anything here will be called only after the next handler's method has finished
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		proto := r.Proto
		method := r.Method
		uri := r.URL.RequestURI()

		app.log.Info("receivedrequest", "ip", ip, "proto", proto, "method", method, "uri", uri)
		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
func (app *application) registerAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			app.clientError(w, r, http.StatusForbidden) // Change to redirect, but make that better
			return
		}
		// Customer claims struct here?
		tkn, err := jwt.ParseWithClaims(c.Value, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return jwtKey, nil
		})

		if err != nil {
			app.clientError(w, r, http.StatusForbidden)
			return
		}

		if !tkn.Valid {
			app.clientError(w, r, http.StatusForbidden)
			return
		}

		claims, ok := tkn.Claims.(*CustomClaims)

		if !ok {
			app.serverError(w, r, fmt.Errorf("couldn't convert token claims field to CustomClaims type"))
			return
		}

		//
		username := claims.Username
		authLevel := claims.AuthLevel
		ctx := r.Context()

		ctx = context.WithValue(ctx, usernameContextKey, username)
		ctx = context.WithValue(ctx, authLevelContextKey, authLevel)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}

func (app *application) requireProvider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authLevel, ok := r.Context().Value(authLevelContextKey).(int)

		if !ok {
			app.clientError(w, r, http.StatusForbidden)
			return
		}

		if authLevel < authLevelProvider {
			app.clientError(w, r, http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authLevel, ok := r.Context().Value(authLevelContextKey).(int)

		if !ok {
			app.clientError(w, r, http.StatusForbidden)
			return
		}

		if authLevel < authLevelProvider {
			app.clientError(w, r, http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
