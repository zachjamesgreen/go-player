package controllers

import (
	"encoding/json"
	"github.com/zachjamesgreen/go-player/models"
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
	albums, err := models.GetAlbumSongs(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(albums)
}

func DeleteAlbum(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	album := models.GetAlbum(id)
	album.Delete()
}
