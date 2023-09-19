package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// AjaxRouteAuthOptional represents an AJAX handler where authentication is optional,
// that returns a response object to be sent as JSON, and a status code.
type AjaxRouteAuthOptional func(
	db *sql.DB,
	userID *uint,
	w http.ResponseWriter,
	r *http.Request,
) (interface{}, int)

// AjaxRouteAuthRequired represents an AJAX handler where authenticaition is mandatory,
// that returns a response object to be sent as JSON, and a status code.
type AjaxRouteAuthRequired func(
	db *sql.DB,
	userID uint,
	w http.ResponseWriter,
	r *http.Request,
) (interface{}, int)

var ajaxHandlersAuthOptional = map[string]map[string]AjaxRouteAuthOptional{
	http.MethodGet: {
		"/ajax/test":       ajaxTest,
		"/ajax/fetchLogin": ajaxFetchLoginHandler,
	},
	http.MethodPost: {
		"/ajax/login": ajaxLoginHandler,
	},
}

var ajaxHandlersAuthRequired = map[string]map[string]AjaxRouteAuthRequired{
	http.MethodGet: {},
	http.MethodPost: {
		"/ajax/logout": ajaxLogoutHandler,
	},
}

func isAjax(r *http.Request) bool {
	return strings.ToLower(r.Header.Get("X-Requested-With")) == "xmlhttprequest"
}

func ajaxHandler(db *sql.DB, userID *uint, w http.ResponseWriter, r *http.Request) {
	// var rt = NewResponseTracker(w)

	var handle = func(handler func() (interface{}, int)) {
		// Verify access to admin routes
		if strings.HasPrefix(r.URL.Path, "/ajax/admin") && (userID == nil || !isAdmin(*userID)) {
			logError(r, userID, fmt.Errorf("forbidden admin access"))
			w.WriteHeader(http.StatusForbidden)
			return
		}
		response, statusCode := handler()
		if statusCode >= 400 {
			w.WriteHeader(statusCode)
			// Send current version stamp
			w.Write([]byte("VersionStamp: " + cacheControlVersionStamp))
			return
		}
		if response != nil {
			js, err := json.Marshal(response)
			if err != nil {
				logError(r, userID, fmt.Errorf("marshalling response: %w", err))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(statusCode) // WriteHeader is called after setting headers
			w.Write(js)
		} else {
			w.WriteHeader(statusCode)
		}
	}

	handlersAuthOptional, foundMethod := ajaxHandlersAuthOptional[r.Method]
	if foundMethod {
		handler, fouundPath := handlersAuthOptional[r.URL.Path]
		if fouundPath {
			handle(func() (interface{}, int) {
				return handler(db, userID, w, r)
			})
			return
		}
	}

	handlersAuthRequired, foundMethod := ajaxHandlersAuthRequired[r.Method]
	if foundMethod {
		handler, fouundPath := handlersAuthRequired[r.URL.Path]
		if fouundPath {
			if userID == nil {
				w.WriteHeader(http.StatusForbidden)
			} else {
				handle(func() (interface{}, int) {
					return handler(db, *userID, w, r)
				})
			}
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func ajaxTest(db *sql.DB, userID *uint, w http.ResponseWriter, r *http.Request) (interface{}, int) {
	var message string
	if userID == nil {
		message = "Message retrieved using AJAX"
	} else {
		message = "Message retrieved using AJAX by authenticated user " + ToString(*userID)
	}
	return struct {
		Message string `json:"message"`
	}{message}, http.StatusOK
}
