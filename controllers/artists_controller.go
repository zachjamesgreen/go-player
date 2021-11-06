package controllers

import (
	"encoding/json"
	"music/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetArtists(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.GetArtists())
}

func GetArtist(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	json.NewEncoder(w).Encode(models.GetArtist(id))
}

func GetArtistSongs(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	json.NewEncoder(w).Encode(models.GetArtistSongs(id))
}

func GetArtistAlbums(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	json.NewEncoder(w).Encode(models.GetArtistAlbums(id))
}

func DeleteArtist(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	artist := models.GetArtist(id)
	artist.Delete()
}
