package controllers

import (
	"encoding/json"
	"music/models"
	"net/http"
)

func GetSongs(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.GetSongs())
}
