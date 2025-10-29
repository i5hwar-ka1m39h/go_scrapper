package main

import "net/http"

func handleError(w http.ResponseWriter, r *http.Request) {
	errorResponse(w, 400, "sending fuck from handle err")
}
