package main

import (
	"net/http"
)

func handleSome(w http.ResponseWriter, r *http.Request) {
	jsonResponseWriter(w, 200, struct{}{})
}
