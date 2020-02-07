package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// AjaxRoute represents an authenticated AJAX handler that returns
// a response object to be sent as JSON, or an error to log, and a status code.
type AjaxRoute func(db *sql.DB, userID uint, w http.ResponseWriter, r *http.Request) (interface{}, int, error)

var ajaxHandlers = map[string]map[string]AjaxRoute{
	http.MethodGet: map[string]AjaxRoute{
		"/ajax/test": ajaxTest,
	},
	http.MethodPost: map[string]AjaxRoute{},
}

func ajaxHandler(db *sql.DB, userID uint, w http.ResponseWriter, r *http.Request) {
	handlers, foundMethod := ajaxHandlers[r.Method]
	if foundMethod {
		handler, fouundPath := handlers[r.URL.Path]
		if fouundPath {
			response, statusCode, err := handler(db, userID, w, r)
			if err != nil {
				log.Printf("Error running ajax handler [%s]: %v\n", r.URL.Path, err)
				w.WriteHeader(statusCode)
				return
			}
			if response != nil {
				js, err := json.Marshal(response)
				if err != nil {
					log.Printf("Error marshalling response: %v\n", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(statusCode)
				w.Write(js)
			}
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func ajaxTest(db *sql.DB, userID uint, w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	response := struct {
		Message string `json:"message"`
	}{"Message retrieved using AJAX"}

	return response, http.StatusOK, nil
}
