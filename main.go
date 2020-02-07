package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/mux"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// TOOD Convert all IDs to int64

type config struct {
	DBUser   string `json:"dbUser"`
	DBPass   string `json:"dbPass"`
	DBName   string `json:"dbName"`
	HTTPPort string `json:"httpPort"`
}

func main() {
	configContents, err := ioutil.ReadFile("env.json")
	if err != nil {
		log.Fatalln(err)
	}

	var config = &config{}
	err = json.Unmarshal(configContents, config)
	if err != nil {
		log.Fatalln(err)
	}

	initDB := flag.Bool("initDB", false, "Initialize a fresh database")
	createNewUser := flag.Bool("createUser", false, "Whether to create a user on startup")
	newUserUsername := flag.String("username", "", "Login username for new user")
	newUserPassword := flag.String("password", "", "Password for new user")
	newUserEmail := flag.String("email", "", "Email address for new user")
	flag.Parse()

	// Include file and line in log output
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db, err := sql.Open("mysql", config.DBUser+":"+config.DBPass+"@/"+config.DBName+"?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	if *initDB {
		// Load initializing SQL
		initFileContents, err := ioutil.ReadFile("sql/init.sql")
		if err != nil {
			log.Fatalln(err)
		}

		// Remove comment lines
		commentMatcher := regexp.MustCompile("(?m)[\r\n]+^--.*$")
		withoutComments := commentMatcher.ReplaceAllString(string(initFileContents), "")

		// Split into statements
		lines := strings.Split(withoutComments, ";")

		// Execute each statement
		for i := 0; i < len(lines); i++ {
			line := strings.TrimSpace(lines[i])
			db.Exec(line)
		}
		log.Println("Database initialized")
	}

	if *createNewUser {
		if _, err := createUser(db, *newUserUsername, *newUserPassword, *newUserEmail); err != nil {
			log.Fatalln(err)
		}
		log.Println("New user created")
	}

	r := mux.NewRouter()

	// set up static resource routes
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	// set up authenticated routes
	authenticate := makeAuthenticator(db)

	r.HandleFunc("/login", makeLoginHandler(db))
	r.HandleFunc("/logout", authenticate(logoutHandler))

	r.PathPrefix("/ajax/").HandlerFunc(authenticate(ajaxHandler))

	// All other paths go through index handler
	r.PathPrefix("/").HandlerFunc(authenticate(indexHandler))

	s := &http.Server{
		Addr:    ":" + config.HTTPPort,
		Handler: r,
	}

	log.Printf("Listening on port %s", config.HTTPPort)

	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

var indexTemplate = template.Must(template.ParseFiles("index.html"))

func indexHandler(db *sql.DB, userID uint, w http.ResponseWriter, r *http.Request) {
	row := db.QueryRow("SELECT username FROM user WHERE id=?", userID)
	var username string
	err := row.Scan(&username)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	indexTemplate.Execute(w, struct {
		UserID   uint
		Username string
	}{userID, username})
}
