package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name: "session",
			Value: uuid.NewV4().String(),
		}
	}
	c.MaxAge = sessionLen
	http.SetCookie(w, c)
	//attribute present and given in seconds
	var u user
	if s, ok := sessions[c.Value]; ok {
		s.ts = time.Now()
		sessions[c.Value] = s
		u = users[s.un]
	}
	return u
}

func isloggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := sessions[c.Value]
	if ok {
		s.ts = time.Now()
		sessions[c.Value] = s
	}
	_, ok = users[s.un]
	//refresh session
	c.MaxAge = sessionLen
	http.SetCookie(w, c)
	return ok
}

func showSessions() {
	fmt.Println("******")
	for k, v := range sessions {
		fmt.Println(k, v.un)
	}
}

func cleanSessions() {
	fmt.Println("*** BEFORE CLEAN ***")
	showSessions()
	for k, v := range sessions {
		if 	time.Since(v.ts) > (time.Second * 30) {
			delete(sessions, k)
		}
	}
	fmt.Println("*** AFTER CLEAN ***")
	showSessions()
}
