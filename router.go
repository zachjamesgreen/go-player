package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"music/models"
)

func mount(r *mux.Router) {
	// --------------
	// Artists Routes
	// --------------
	r.HandleFunc("/artists", func(w http.ResponseWriter, res *http.Request) {
		var artists = models.GetArtists()
		json.NewEncoder(w).Encode(artists)
	})
	r.HandleFunc("/artists/{id}", func(w http.ResponseWriter, res *http.Request) {
		vars := mux.Vars(res)
		id, err := strconv.Atoi(vars["id"])
		check(err)
		var artist = models.GetArtist(id)
		json.NewEncoder(w).Encode(artist)
	})
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
