package main

import (
	"github.com/zachjamesgreen/go-player/controllers"

	"github.com/gorilla/mux"
)

func mount(r *mux.Router) {
	// --------------
	// Artists Routes
	// --------------
	r.HandleFunc("/artists", controllers.GetArtists)
	r.HandleFunc("/artists/{id}", controllers.DeleteArtist).Methods("DELETE")
	r.HandleFunc("/artists/{id}", controllers.GetArtist)
	r.HandleFunc("/artists/{id}/songs", controllers.GetArtistSongs)
	r.HandleFunc("/artists/{id}/albums", controllers.GetArtistAlbums)
	// ------------
	// Album Routes
	// ------------
	r.HandleFunc("/albums", controllers.GetAlbums)
	r.HandleFunc("/albums/{id}", controllers.DeleteAlbum).Methods("DELETE")
	r.HandleFunc("/albums/{id}", controllers.GetAlbum)
	r.HandleFunc("/albums/{id}/songs", controllers.GetAlbumSongs)
	// -----------
	// Song Routes
	// -----------
	r.HandleFunc("/songs", controllers.GetSongs)
	r.HandleFunc("/songs/liked/{id}", controllers.LikeSong)
	r.HandleFunc("/songs/liked", controllers.GetLikedSongs)
	r.HandleFunc("/songs/liked/{id}/remove", controllers.UnlikeSong)
	r.HandleFunc("/songs/{id}", controllers.DeleteSong).Methods("DELETE")

	// ------------
	// Upload Route
	// ------------
	r.HandleFunc("/upload", UploadHandler) //.Methods("POST")

	// -----------
	// User Routes
	// -----------
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.GetUserById)

	//------------------
	// Liked Songs Route
	//------------------
	// r.HandleFunc("/liked/{id}/remove", UnlikeSong)
	// r.HandleFunc("/liked/{id}", LikeSong)
	// r.HandleFunc("/liked", GetLikedSongs)

}
