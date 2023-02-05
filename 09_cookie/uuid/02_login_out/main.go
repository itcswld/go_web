package main

import (
	"html/template"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username string
	Password []byte
	First string
	Last string
	Age string
}
type session struct {
	un string
	ts time.Time
}

var tpl *template.Template
var users = map[string]user{}
var sessions = map[string]session{}
var now time.Time
const sessionLen int = 30

func init() {
	tpl = template.Must(template.ParseGlob("tpl/*"))
	un := "admin@test.com"
	p := "password"
	bs, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	users[un] = user{un, bs, "Eve", "Tong", "adult"}
	now = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	showSessions()
	tpl.ExecuteTemplate(w, "index.html", u)
}
func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !isloggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Age != "adult" {
		http.Error(w, "You must be adult to enter the bar",http.StatusForbidden)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "bar.html", u)
}
//signup
func signup(w http.ResponseWriter, req *http.Request) {
	if isloggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		var u user
		un := req.FormValue("un")
		p := req.FormValue("pwd")
		f := req.FormValue("fn")
		l := req.FormValue("ln")
		a := req.FormValue("age")

		if _, ok := users[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		uuID := uuid.NewV4()
		c := &http.Cookie{
			Name: "session",
			Value: uuID.String(),
		}
		http.SetCookie(w, c)
		sessions[c.Value] = session{un, time.Now()}
		bc, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		u = user{un, bc, f, l, a}
		users[un] = u
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "signup.html", nil)
}
//login
func login(w http.ResponseWriter, req *http.Request) {
	if isloggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		un := req.FormValue("un")
		p := req.FormValue("pwd")
		//is there a username ? or does the pwd match the stored pwd ?
		u, ok := users[un]
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if !ok || err != nil {
			http.Error(w, "Username and/or password does not match", http.StatusForbidden)
			return
		}
		//create session
		c := &http.Cookie{
			Name: "session",
			Value: uuid.NewV4().String(),
			MaxAge: sessionLen,
		}
		http.SetCookie(w, c)
		sessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "login.html", nil)
}
//logout
func logout(w http.ResponseWriter, req *http.Request) {
	if !isloggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//get cookie
	c, _ := req.Cookie("session")
	//delete
	delete(sessions, c.Value)
	c = &http.Cookie {
		Name: "session",
		Value: "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	//clean up sessions
	lastTimeLoggedin := sessions[c.Value].ts
	if 	time.Since(lastTimeLoggedin) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
