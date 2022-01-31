package main

import (
	
	"net/http"

	"github.com/justinas/nosurf"
)


// NoSurf is the csrf protection middleware
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

//LoadAndSave provides middleware which automatically loads and saves session data for the current request, and communicates the session token to and from the client in a cookie.
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
