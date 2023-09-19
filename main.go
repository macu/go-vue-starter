package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"           // http routing (BSD-3-Clause)
	_ "github.com/jackc/pgx/v4/stdlib" // postgres (MIT)
)

type config struct {
	DBUser       string `json:"dbUser"`
	DBPass       string `json:"dbPass"`
	DBName       string `json:"dbName"`
	HTTPPort     uint   `json:"httpPort"`
	Debug        bool   `json:"debug"`
	VersionStamp string `json:"versionStamp"`
}

var cacheControlVersionStamp string
var debugMode bool

var indexTemplate = template.Must(template.ParseFiles("html/index.html"))

func main() {

	if StringToBool(os.Getenv("MAINTENANCE_MODE")) {
		// Site down for databae upgrades
		maintenanceMode()
		os.Exit(0)
	}

	configContents, err := os.ReadFile("env.json")
	if err != nil {
		log.Fatalln(err)
	}

	var config = &config{}
	err = json.Unmarshal(configContents, config)
	if err != nil {
		log.Fatalln(err)
	}

	debugMode = config.Debug
	cacheControlVersionStamp = config.VersionStamp

	var createNewUser = flag.Bool("createUser", false, "Whether to create a user on startup")
	var newUserUsername = flag.String("username", "", "Login username for new user")
	var newUserPassword = flag.String("password", "", "Password for new user")
	var newUserEmail = flag.String("email", "", "Email address for new user")
	flag.Parse()

	// Include file and line in log output
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// postgres:
	var dataSource = fmt.Sprintf("host=%s user=%s password=%s database=%s",
		"localhost", config.DBUser, config.DBPass, config.DBName)

	db, err := sql.Open("pgx", dataSource)
	if err != nil {
		logErrorFatal(err)
	}

	if *createNewUser {
		if _, err := createUser(new(http.Request),
			db, *newUserUsername, *newUserPassword, *newUserEmail); err != nil {
			log.Fatalln(err)
		}
		log.Println("New user created")
	}

	var r = mux.NewRouter()

	// set up static resource routes
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	// set up authenticated routes
	authenticate := makeAuthenticator(db)

	r.PathPrefix("/ajax/").HandlerFunc(authenticate(ajaxHandler))

	// All other paths go through index handler
	r.PathPrefix("/").HandlerFunc(indexHandler)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.HTTPPort),
		Handler: r,
	}

	log.Printf("Listening on port %d", config.HTTPPort)

	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, struct {
		VersionStamp string
		Debug        bool
	}{cacheControlVersionStamp, debugMode})
}

func maintenanceMode() {
	var port string

	var config = &config{}
	configContents, err := os.ReadFile("env.json")
	if err != nil {
		logErrorFatal(err)
	}
	err = json.Unmarshal(configContents, config)
	if err != nil {
		logErrorFatal(err)
	}
	port = ToString(config.HTTPPort)

	var maintenancePageTemplate *template.Template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		if !isAjax(r) {
			if maintenancePageTemplate == nil {
				maintenancePageTemplate = template.Must(template.ParseFiles("html/maintenance.html"))
			}
			maintenancePageTemplate.Execute(w, nil)
		}
	})

	log.Printf("Listening on port %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != http.ErrServerClosed {
		logErrorFatal(err)
	}
}
