package main

import (
	"encoding/json"
	"music/models"
	"net/http"
)

func GetSongs(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(models.GetSongs())
}
