package main

import (
	"net/http"
)

func Server() {

	http.ListenAndServe(":5000", getRoutes())
}

func getRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/annonces", GetAllAnnnoncesHandler)
	return mux

}
