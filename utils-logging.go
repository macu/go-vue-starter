package main

import (
	"log"
	"net/http"
	"runtime"
)

func logDefault(r *http.Request, data interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	log.Printf("[default] %s:%d %v", fn, line, data)
}

func logNotice(r *http.Request, data interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	log.Printf("[notice] %s:%d %v", fn, line, data)
}

func logError(r *http.Request, userID *uint, err error) {
	if err == nil {
		return
	}

	_, fn, line, _ := runtime.Caller(1)
	log.Printf("[error] %s:%d %+v", fn, line, err)
}

func logErrorFatal(err error) {
	// Continue even if err is nil

	_, fn, line, _ := runtime.Caller(1)
	log.Fatalf("[fatal] %s:%d %+v", fn, line, err)
}
