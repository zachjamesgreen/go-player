package main

import (
	"encoding/json"
	"music/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAlbums(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.GetAlbums())
}

func GetAlbum(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	json.NewEncoder(w).Encode(models.GetAlbum(id))
}

func GetAlbumSongs(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	json.NewEncoder(w).Encode(models.GetAlbumSongs(id))
}
