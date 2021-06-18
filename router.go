package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func mount(r *mux.Router) {
	// --------------
	// Artists Routes
	// --------------
	r.HandleFunc("/artists", func(w http.ResponseWriter, res *http.Request) {
		var artists []Artist = getArtists()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(artists)
	})
	r.HandleFunc("/artists/{id}", func(w http.ResponseWriter, res *http.Request) {
		vars := mux.Vars(res)
		id, err := strconv.Atoi(vars["id"])
		check(err)
		var artist Artist = getArtist(id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(artist)
	})
	r.HandleFunc("/artists/{id}/songs", func(w http.ResponseWriter, res *http.Request) {
		vars := mux.Vars(res)
		id, err := strconv.Atoi(vars["id"])
		check(err)
		songs := getArtistSongs(id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(songs)
	})
	r.HandleFunc("/artists/{id}/albums", func(w http.ResponseWriter, res *http.Request) {
		vars := mux.Vars(res)
		id, err := strconv.Atoi(vars["id"])
		check(err)
		songs := getArtistAlbums(id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(songs)
	})
	// ------------
	// Album Routes
	// ------------
	r.HandleFunc("/albums", func(w http.ResponseWriter, res *http.Request) {
		albums := getAlbums()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(albums)
	})
	r.HandleFunc("/albums/{id}", func(w http.ResponseWriter, res *http.Request) {
		vars := mux.Vars(res)
		id, err := strconv.Atoi(vars["id"])
		check(err)
		album := getAlbum(id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(album)
	})
	r.HandleFunc("/albums/{id}/songs", func(w http.ResponseWriter, res *http.Request) {
		vars := mux.Vars(res)
		id, err := strconv.Atoi(vars["id"])
		check(err)
		songs := getAlbumSongs(id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(songs)
	})
	// -----------
	// Song Routes
	// -----------
	r.HandleFunc("/songs", func(w http.ResponseWriter, res *http.Request) {
		songs := getSongs()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(songs)
	})

	// ------------
	// Upload Route
	// ------------
	r.HandleFunc("/upload", func(w http.ResponseWriter, res *http.Request) {
		res.ParseMultipartForm(10 << 20)
		file, handler, err := res.FormFile("song")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		defer file.Close()
		upload(file, handler)
		// return that we have successfully uploaded our file!
		fmt.Fprintf(w, "Successfully Uploaded File\n")

	}).Methods("POST")
}
