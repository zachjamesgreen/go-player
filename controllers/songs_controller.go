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
	json.NewEncoder(w).Encode(models.GetSongs())
}

func GetLikedSongs(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	songs := models.GetLikedSongs()
	json.NewEncoder(w).Encode(songs)
}

func LikeSong(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	json.NewEncoder(w).Encode(models.AddLike(id))
}

func UnlikeSong(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	json.NewEncoder(w).Encode(models.RemoveLike(id))
}

func DeleteSong(w http.ResponseWriter, req *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	song := models.GetSong(id)
	song.Delete()
	// json.NewEncoder(w).Encode()
}
