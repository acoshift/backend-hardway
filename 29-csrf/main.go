package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signin", signIn)
	http.HandleFunc("/signout", signOut)
	http.HandleFunc("/transfer", transfer)
	http.ListenAndServe(":3333", nil)
}

var sessionStore = make(map[string]*session)

type session struct {
	ID     string
	UserID int
	// add csrf token
}

func findSession(sessionID string) *session {
	return sessionStore[sessionID]
}

func storeSession(sessionID string, sess *session) {
	sessionStore[sessionID] = sess
}

func getSessionFromRequest(w http.ResponseWriter, r *http.Request) *session {
	var sess *session

	if c, err := r.Cookie("session"); err == nil {
		sess = findSession(c.Value)
	}

	if sess == nil {
		sess = &session{
			ID: generateSessionID(),
		}
		storeSession(sess.ID, sess)
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    sess.ID,
			Path:     "/",
			HttpOnly: true,
		})
	}
	return sess
}

func generateSessionID() string {
	b := make([]byte, 16)
	io.ReadFull(rand.Reader, b)
	return base64.RawURLEncoding.EncodeToString(b)
}

func generateCSRFToken() string {
	//
	return ""
}

func index(w http.ResponseWriter, r *http.Request) {
	var userID int

	sess := getSessionFromRequest(w, r)
	if sess != nil {
		userID = sess.UserID
	}

	if userID == 0 {
		w.Write([]byte(`
			<!doctype html>
			<a href=/signin>Sign In</a>
		`))
		return
	}

	// add csrf token to form
	fmt.Fprintf(w, `
		<!doctype html>
		<form method=POST action=/transfer>
			<input name=amount placeholder=amount required>
			<button type=submit>Transfer</button>
		</form>
		<a href=/signout>Sign Out</a>
	`)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	// session fixation
	sess := getSessionFromRequest(w, r)
	sess.UserID = 1

	// generate new csrf token

	http.Redirect(w, r, "/", http.StatusFound)
}

func signOut(w http.ResponseWriter, r *http.Request) {
	sess := getSessionFromRequest(w, r)
	sess.UserID = 0

	http.Redirect(w, r, "/", http.StatusFound)
}

func transfer(w http.ResponseWriter, r *http.Request) {
	// allow only POST

	// check origin

	// check referer

	// check is user sign in

	// get amount from form

	// get csrf token from form

	// check form csrf token with session csrf token

	// if transfer success print amount to console

	// redirect to /
}
