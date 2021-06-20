package main

import (
	"encoding/json"
	"music/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetArtists(w http.ResponseWriter, res *http.Request) {
	json.NewEncoder(w).Encode(models.GetArtists())
}

func GetArtist(w http.ResponseWriter, res *http.Request) {
	vars := mux.Vars(res)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	json.NewEncoder(w).Encode(models.GetArtist(id))
}
