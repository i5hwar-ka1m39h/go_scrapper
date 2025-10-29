package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func errorResponse(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("something went wrong", message)
	}
	type errResp struct {
		Error string `json:"error"`
	}

	jsonResponseWriter(w, code, errResp{
		Error: message,
	})
}

func jsonResponseWriter(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("error occured in json reponse writer", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
