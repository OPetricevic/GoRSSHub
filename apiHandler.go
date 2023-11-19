package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct{}{})
}

func handleError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 200, "Something went wrong")
}
