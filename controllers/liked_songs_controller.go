package controllers

import (
	"encoding/json"
	"log"
	"music/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var ls models.LikedSong

func GetLikedSongs(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := ls.All()
	check(err)
	json.NewEncoder(w).Encode(rows)
}

func LikeSong(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	err = ls.Add(id)
	if err != nil {
		log.Println(err)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func UnlikeSong(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	check(err)
	err = ls.Remove(id)
	if err != nil {
		log.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
