package main

import (
	"github.com/gorilla/mux"
)

func mount(r *mux.Router) {
	// --------------
	// Artists Routes
	// --------------
	r.HandleFunc("/artists", GetArtists)
	r.HandleFunc("/artists/{id}", GetArtist)
	r.HandleFunc("/artists/{id}/songs", GetArtistSongs)
	r.HandleFunc("/artists/{id}/albums", GetArtistAlbums)
	// ------------
	// Album Routes
	// ------------
	r.HandleFunc("/albums", GetAlbums)
	r.HandleFunc("/albums/{id}", GetAlbum)
	r.HandleFunc("/albums/{id}/songs", GetAlbumSongs)
	// -----------
	// Song Routes
	// -----------
	r.HandleFunc("/songs", GetSongs)

	// ------------
	// Upload Route
	// ------------
	r.HandleFunc("/upload", UploadHandler).Methods("POST")
}
