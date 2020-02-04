package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

var ajaxHandlers = map[string]map[string]AuthenticatedRoute{
	http.MethodGet: map[string]AuthenticatedRoute{
		"/ajax/test": ajaxTest,
	},
	http.MethodPost: map[string]AuthenticatedRoute{},
}

func ajaxHandler(db *sql.DB, userID uint, w http.ResponseWriter, r *http.Request) {
	handlers, foundMethod := ajaxHandlers[r.Method]
	if foundMethod {
		handler, fouundPath := handlers[r.URL.Path]
		if fouundPath {
			handler(db, userID, w, r)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func ajaxTest(db *sql.DB, userID uint, w http.ResponseWriter, r *http.Request) {
	response := struct {
		Message string `json:"message"`
	}{"Message retrieved using AJAX"}

	js, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
