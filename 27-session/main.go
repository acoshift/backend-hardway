package main

import (
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
	//
}

func findSession(sessionID string) *session {
	return nil
}

func storeSession(sessionID string, sess *session) {
	//
}

func getSessionFromRequest(w http.ResponseWriter, r *http.Request) *session {
	// read session id from cookie

	// find session from database

	// if session not found create new session, and set session id to cookie

	return nil
}

func generateSessionID() string {
	// create new slice

	// read data from rand.Reader from package crypto/rand

	// encode to base64

	return ""
}

func index(w http.ResponseWriter, r *http.Request) {
	// get session from request

	// if not sign in, return sign in page

	// if already sign in return sign out page
}

func signIn(w http.ResponseWriter, r *http.Request) {
	// TODO: session fixation hack

	// get session from request

	// set user id to `1`

	// redirect to /
}

func signOut(w http.ResponseWriter, r *http.Request) {
	// get session from request

	// set user id to `0`

	// redirect to /
}
