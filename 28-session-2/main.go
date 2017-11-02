package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signin", signIn)
	http.HandleFunc("/signout", signOut)
	http.ListenAndServe(":3333", nil)
}

var sessionStore = make(map[string]*session)

type session struct {
	ID     string
	UserID int
}

func findSession(sessionID string) *session {
	return sessionStore[sessionID]
}

func storeSession(sessionID string, sess *session) {
	sessionStore[sessionID] = sess
}

func removeSession(sessionID string) {
	//
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

func rotateSession(w http.ResponseWriter, sess *session) *session {
	// copy old session data to new session

	// generate new session id

	// save new session to storage

	// remove old Set-Cookie header

	// add session id to cookie

	// remove old session from storage

	// return new session
	return nil
}

func generateSessionID() string {
	b := make([]byte, 16)
	io.ReadFull(rand.Reader, b)
	return base64.RawURLEncoding.EncodeToString(b)
}

func index(w http.ResponseWriter, r *http.Request) {
	var userID int

	sess := getSessionFromRequest(w, r)
	if sess != nil {
		userID = sess.UserID
	}

	// not sign in
	if userID == 0 {
		w.Write([]byte(`
			<!doctype html>
			<a href=/signin>Sign In</a>
		`))
		return
	}

	w.Write([]byte(`
		<!doctype html>
		<a href=/signout>Sign Out</a>
	`))
}

func signIn(w http.ResponseWriter, r *http.Request) {
	sess := getSessionFromRequest(w, r)

	// rotate session id

	sess.UserID = 1

	http.Redirect(w, r, "/", http.StatusFound)
}

func signOut(w http.ResponseWriter, r *http.Request) {
	sess := getSessionFromRequest(w, r)
	sess.UserID = 0

	http.Redirect(w, r, "/", http.StatusFound)
}
