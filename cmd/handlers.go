package main

import (
	"encoding/json"
	"net/http"
)

func GetAllAnnnoncesHandler(w http.ResponseWriter, r *http.Request) {

	annonces := GetAllAnnnonces()

	bytes, _ := json.Marshal(annonces)
	w.Write(bytes)

}
