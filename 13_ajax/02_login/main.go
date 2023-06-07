package main

import (
	"fmt"
	"html/template"
	"io/ioutil" // I/O utility
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	userName     string
	lastActivity time.Time
}

var tpl *template.Template
var users = map[string]user{}       //map[session.username]
var sessions = map[string]session{} //map[uid]
var lastTimeCleaned time.Time

const sessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("tpl/*"))
	lastTimeCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	//ajax route
	http.HandleFunc("/checkUserName", checkUserName)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	showSession()
	tpl.ExecuteTemplate(w, "index.html", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "you must be 007 to enter the bar.", http.StatusForbidden)
		return
	}
	showSession()
	tpl.ExecuteTemplate(w, "bar.html", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//onPress Submission
	var n string
	if req.Method == http.MethodPost {
		fmt.Println("got post values~")
		//get form values
		n = req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		//username taken?
		if _, ok := users[n]; ok {
			http.Error(w, "Username already taken.", http.StatusForbidden)
			return
		}
		//cookie: create session
		c := &http.Cookie{
			Name:  "session",
			Value: uuid.NewV4().String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		//map: users & sessions storing
		sessions[c.Value] = session{n, time.Now()}
		//crypt password
		cryptedPwd, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Interal server error.", http.StatusInternalServerError)
			return
		}
		users[n] = user{n, cryptedPwd, f, l, r}
		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSession()
	tpl.ExecuteTemplate(w, "signup.html", users[n])
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var n string
	if req.Method == http.MethodPost {
		n = req.FormValue("username")
		p := req.FormValue("password")
		///match?
		//username match?
		if _, ok := users[n]; !ok {
			http.Error(w, "Username and/or password do not match.", http.StatusForbidden)
			return
		}
		//password match?
		err := bcrypt.CompareHashAndPassword(users[n].Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match.", http.StatusForbidden)
			return
		}
		//cookie: create session
		c := &http.Cookie{
			Name:  "session",
			Value: uuid.NewV4().String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		fmt.Print("before login")
		showSession()
		sessions[c.Value] = session{users[n].UserName, time.Now()}
		fmt.Print("after login")
		showSession()
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSession()
	tpl.ExecuteTemplate(w, "login.html", users[n])
}

func logout(w http.ResponseWriter, req *http.Request) {
	c, _ := req.Cookie("session")
	//delete the session
	delete(sessions, c.Value)
	//remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	//timeout ? clean all session
	if time.Since(lastTimeCleaned) > (time.Second * 30) {
		go cleanSessions()
	}
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

//ajax route
func checkUserName(w http.ResponseWriter, req *http.Request) {
	sampleUsers := map[string]bool{
		"eve@eve.com":     true,
		"love@love.com":   true,
		"money@money.com": true,
	}

	userName, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	n := string(userName)
	fmt.Println("Username:", n)

	fmt.Fprint(w, sampleUsers[n])
}
