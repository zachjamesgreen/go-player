package main

import (
	. "github.com/zachjamesgreen/go-player/controllers"

	"github.com/gorilla/mux"
)

func mount(r *mux.Router) {
	// --------------
	// Artists Routes
	// --------------
	r.HandleFunc("/artists", GetArtists)
	r.HandleFunc("/artists/{id}", DeleteArtist).Methods("DELETE")
	r.HandleFunc("/artists/{id}", GetArtist)
	r.HandleFunc("/artists/{id}/songs", GetArtistSongs)
	r.HandleFunc("/artists/{id}/albums", GetArtistAlbums)
	// ------------
	// Album Routes
	// ------------
	r.HandleFunc("/albums", GetAlbums)
	r.HandleFunc("/albums/{id}", DeleteAlbum).Methods("DELETE")
	r.HandleFunc("/albums/{id}", GetAlbum)
	r.HandleFunc("/albums/{id}/songs", GetAlbumSongs)
	// -----------
	// Song Routes
	// -----------
	r.HandleFunc("/songs", GetSongs)
	r.HandleFunc("/songs/liked/{id}", LikeSong)
	r.HandleFunc("/songs/liked", GetLikedSongs)
	r.HandleFunc("/songs/liked/{id}/remove", UnlikeSong)
	r.HandleFunc("/songs/{id}", DeleteSong).Methods("DELETE")

	// ------------
	// Upload Route
	// ------------
	r.HandleFunc("/upload", UploadHandler) //.Methods("POST")

	// -----------
	// User Routes
	// -----------
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", GetUserById)

	//------------------
	// Liked Songs Route
	//------------------
	// r.HandleFunc("/liked/{id}/remove", UnlikeSong)
	// r.HandleFunc("/liked/{id}", LikeSong)
	// r.HandleFunc("/liked", GetLikedSongs)

}
