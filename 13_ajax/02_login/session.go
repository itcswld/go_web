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
	//the session is not exists? create cookie
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: uuid.NewV4().String(),
		}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	//the user exists already? get user
	var u user
	if s, ok := sessions[c.Value]; ok {
		//update timestamp
		s.lastActivity = time.Now()
		sessions[c.Value] = s
		u = users[s.userName]
	}
	return u
}

func showSession() {
	fmt.Println("****************")
	for k, v := range sessions {
		fmt.Println(k, v.userName)
	}
	fmt.Println("")
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := sessions[c.Value]
	if ok {
		//update timestamp
		s.lastActivity = time.Now()
		sessions[c.Value] = s
	}
	_, ok = users[s.userName]
	c.MaxAge = sessionLength
	//renew cookie
	http.SetCookie(w, c)

	return ok
}

func cleanSessions() {
	fmt.Println("*** before clean session***")
	showSession()
	for k, v := range sessions {
		fmt.Println(time.Since(v.lastActivity))
		//timeout?
		if time.Since(v.lastActivity) > (time.Second * 30) {
			delete(sessions, k)
		}
		lastTimeCleaned = time.Now()
		fmt.Println("***after clean session***")
		showSession()
	}
}
