package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func (h *Handler) handleRequest(hp HandlerParam) {
	*h.c++
	var err error = nil
	defer func() {
		apiLog(h.l, h.c, &hp.r.RequestURI, err)
	}()

	err = checkHTTPMethod(hp.w, hp.r.Method, hp.method)
	if err != nil {
		return
	}

	err = hp.handlerFunc(hp.w, hp.r)
	if err != nil {
		http.Error(hp.w, err.Error(), http.StatusInternalServerError)
	}
}

func apiLog(l *log.Logger, counter *uint, url *string, err error) {
	var status string
	if err == nil {
		status = "SUCCESS"
	} else {
		status = err.Error()
	}

	l.Printf("[%d] [%s] [%s]", *counter, *url, status)
}

func toJSON(w http.ResponseWriter, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func checkHTTPMethod(w http.ResponseWriter, reqMethod, desMethod string) error {
	if reqMethod != desMethod {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return errors.New("invalid http method")
	}
	return nil
}
