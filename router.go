package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"music/models"

	"github.com/gorilla/mux"
)

func mount(r *mux.Router) {
	// --------------
	// Artists Routes
	// --------------
	r.HandleFunc("/artists", GetArtists)
	r.HandleFunc("/artists/{id}", GetArtist)
	r.HandleFunc("/artists/{id}/songs", func(w http.ResponseWriter, res *http.Request) {
		vars := mux.Vars(res)
		id, err := strconv.Atoi(vars["id"])
		check(err)
		songs := models.GetArtistSongs(id)
		json.NewEncoder(w).Encode(songs)
	})
	r.HandleFunc("/artists/{id}/albums", func(w http.ResponseWriter, res *http.Request) {
		vars := mux.Vars(res)
		id, err := strconv.Atoi(vars["id"])
		check(err)
		songs := models.GetArtistAlbums(id)
		json.NewEncoder(w).Encode(songs)
	})
	// ------------
	// Album Routes
	// ------------
	r.HandleFunc("/albums", func(w http.ResponseWriter, res *http.Request) {
		albums := models.GetAlbums()
		fmt.Print(albums)
		json.NewEncoder(w).Encode(albums)
	})
	r.HandleFunc("/albums/{id}", func(w http.ResponseWriter, res *http.Request) {
		vars := mux.Vars(res)
		id, err := strconv.Atoi(vars["id"])
		check(err)
		album := models.GetAlbum(id)
		json.NewEncoder(w).Encode(album)
	})
	r.HandleFunc("/albums/{id}/songs", func(w http.ResponseWriter, res *http.Request) {
		vars := mux.Vars(res)
		id, err := strconv.Atoi(vars["id"])
		check(err)
		songs := models.GetAlbumSongs(id)
		json.NewEncoder(w).Encode(songs)
	})
	// -----------
	// Song Routes
	// -----------
	r.HandleFunc("/songs", func(w http.ResponseWriter, res *http.Request) {
		songs := models.GetSongs()
		// fmt.Print(songs)
		json.NewEncoder(w).Encode(songs)
	})

	// ------------
	// Upload Route
	// ------------
	r.HandleFunc("/upload", func(w http.ResponseWriter, req *http.Request) {
		req.ParseMultipartForm(32 << 20)
		files := req.MultipartForm.File["song"]
		fmt.Println(files)

		for _, fileHeader := range files {
			fmt.Printf("Uploaded File: %+v\n", fileHeader.Filename)
			fmt.Printf("File Size: %+v\n", fileHeader.Size)
			fmt.Printf("MIME Header: %+v\n", fileHeader.Header)
			upload(fileHeader)
		}

		// return that we have successfully uploaded our file!
		fmt.Fprintf(w, "Successfully Uploaded File\n")

	}).Methods("POST")
}
