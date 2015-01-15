package main

import (
	"database/sql"
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/codegangsta/negroni"
	"golang.org/x/crypto/bcrypt"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	_ "github.com/go-sql-driver/mysql"
	"github.com/unrolled/render"
	"fmt"
	"log"
	"net/http"
)

type tomlConfig struct {
	DB database `toml:"database"`
}

type database struct {
	User string
	Password string
	DBName string
}

var db *sql.DB = setupDB()

func main() {

	defer db.Close()

	mux := http.NewServeMux()
	n := negroni.Classic()

	store := cookiestore.New([]byte("secretkey789"))
	n.Use(sessions.Sessions("global_session_store", store))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		SimplePage(w, r, "mainpage")
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			SimplePage(w, r, "login")
		} else if r.Method == "POST" {
			LoginPost(w, r)
		}
	})

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			SimplePage(w, r, "signup")
		} else if r.Method == "POST" {
			SignupPost(w, r)
		}
	})

	mux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		Logout(w, r)
	})

	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		SimpleAuthenticatedPage(w, r, "home")
	})

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		APIHandler(w, r)
	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	n.UseHandler(mux)
	n.Run(":8080")
	fmt.Println("Server running on port 8080")

}

func setupDB() *sql.DB {

	var config tomlConfig
		if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Print(err)
	}

	db, err := sql.Open("mysql", config.DB.User + ":" + config.DB.Password + "@/" + config.DB.DBName + "?charset=utf8")
	if err != nil {
		panic(err)
	}

	return db

}


func errHandler(err error) {
	if err != nil {
		log.Print(err)
	}
}


func SimplePage(w http.ResponseWriter, req *http.Request, template string) {

	r := render.New(render.Options{})
	r.HTML(w, http.StatusOK, template, nil)

}


func SimpleAuthenticatedPage(w http.ResponseWriter, req *http.Request, template string) {

	session := sessions.GetSession(req)
	sess := session.Get("useremail")

	if sess == nil {
		http.Redirect(w, req, "/notauthenticated", 301)
	}

	r := render.New(render.Options{})
	r.HTML(w, http.StatusOK, template, nil)

}

func LoginPost(w http.ResponseWriter, req *http.Request) {

	session := sessions.GetSession(req)

	username := req.FormValue("inputUsername")
	password := req.FormValue("inputPassword")

	var (
		email string
		hashed_password string
	)

    err := db.QueryRow("SELECT user_email, user_password FROM users WHERE user_name = ? AND user_password = ?", username, password).Scan(&email, &hashed_password)
	password_err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	
	if err != nil && password_err != nil {
		log.Print(err)
		log.Print(password_err)
		http.Redirect(w, req, "/authfail", 301)
	}

	session.Set("useremail", email)
	http.Redirect(w, req, "/home", 302)

}

func SignupPost(w http.ResponseWriter, req *http.Request) {

	username := req.FormValue("inputUsername")
	password := req.FormValue("inputPassword")
	email := req.FormValue("inputEmail")
	hashed_password,hash_err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hash_err !=nil{
		log.Print(hash_err)
	}
	_, err := db.Exec("INSERT INTO users (user_name, user_password, user_email) VALUES (?, ?, ?)", username, hashed_password, email)
	if err != nil {
		log.Print(err)
	}

	http.Redirect(w, req, "/login", 302)

}


func Logout(w http.ResponseWriter, req *http.Request) {

    session := sessions.GetSession(req)
    session.Delete("useremail") 
    http.Redirect(w, req, "/", 302)

}

func APIHandler(w http.ResponseWriter, req *http.Request) {

	data, _ := json.Marshal("{'API Test':'Works!'}")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(data)

}
