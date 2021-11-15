package controllers

import (
	"encoding/json"
	"github.com/zachjamesgreen/go-player/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetSongs(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	songs, err := models.GetSongs()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(songs)
}

func GetLikedSongs(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	songs, err := models.GetLikedSongs()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(songs)
}

func LikeSong(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	song, err := models.GetSong(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	song.AddLike()
	json.NewEncoder(w).Encode(song)
}

func UnlikeSong(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	song, err := models.GetSong(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	song.RemoveLike()
	json.NewEncoder(w).Encode(song)
}

func DeleteSong(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	song, err := models.GetSong(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	song.Delete()
}
