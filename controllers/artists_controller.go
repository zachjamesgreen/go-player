package controllers

import (
	"encoding/json"
	"github.com/zachjamesgreen/go-player/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetArtists(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	artists, err := models.GetAllArtists()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(artists)
}

func GetArtist(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	artist, err := models.GetArtistById(id)
	// handle artist not found. send 404
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(artist)
}

func GetArtistSongs(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	songs, err := models.GetArtistSongsById(id)
	// handle artist not found. send 404
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(songs)
}

func GetArtistAlbums(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	albums, err := models.GetArtistAlbumsById(id)
	// handle artist not found. send 404
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(albums)
}

func DeleteArtist(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	artist, err := models.GetArtistById(id)
	if artist.Name == "" {
		http.Error(w, "Artist Not Found", http.StatusNotFound)
		return
	}
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if err = artist.Delete(); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
}
