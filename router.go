package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/dhowden/tag"
	"github.com/gorilla/mux"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func mount(r *mux.Router, db *sql.DB) {
	// --------------
	// Artists Routes
	// --------------
	r.HandleFunc("/artists", func(w http.ResponseWriter, res *http.Request) {
		var artists []Artist = getArtists(db)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(artists)
	})
	r.HandleFunc("/artists/{id}", func(w http.ResponseWriter, res *http.Request) {
		vars := mux.Vars(res)
		id, err := strconv.Atoi(vars["id"])
		check(err)
		var artist Artist = getArtist(db, id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(artist)
	})
	r.HandleFunc("/artists/{id}/songs", func(w http.ResponseWriter, res *http.Request) {
		vars := mux.Vars(res)
		id, err := strconv.Atoi(vars["id"])
		check(err)
		songs := getArtistSongs(db, id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(songs)
	})
	r.HandleFunc("/artists/{id}/albums", func(w http.ResponseWriter, res *http.Request) {
		vars := mux.Vars(res)
		id, err := strconv.Atoi(vars["id"])
		check(err)
		songs := getArtistAlbums(db, id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(songs)
	})
	// ------------
	// Album Routes
	// ------------
	r.HandleFunc("/albums", func(w http.ResponseWriter, res *http.Request) {
		albums := getAlbums(db)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(albums)
	})
	r.HandleFunc("/albums/{id}", func(w http.ResponseWriter, res *http.Request) {
		vars := mux.Vars(res)
		id, err := strconv.Atoi(vars["id"])
		check(err)
		album := getAlbum(db, id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(album)
	})
	r.HandleFunc("/albums/{id}/songs", func(w http.ResponseWriter, res *http.Request) {
		vars := mux.Vars(res)
		id, err := strconv.Atoi(vars["id"])
		check(err)
		songs := getAlbumSongs(db, id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(songs)
	})
	// -----------
	// Song Routes
	// -----------
	r.HandleFunc("/songs", func(w http.ResponseWriter, res *http.Request) {
		songs := getSongs(db)
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
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)
		// Get tags
		data, err := tag.ReadFrom(file)
		if err != nil {
			log.Fatal(err)
		}
		var artist Artist = Artist{Name: data.Artist()}
		var album Album = Album{Title: data.Album(), ArtistId: 0}
		track, _ := data.Track()
		var song Song = Song{Title: data.Title(), Track: track, Comment: data.Comment(), Genre: data.Genre(), ArtistId: 0, AlbumId: 0}
		var artist_id string
		var album_id string
		var genre string
		ar := db.QueryRow("INSERT INTO artists (name) VALUES ($1) ON CONFLICT (name) DO UPDATE SET name=EXCLUDED.name returning id;", artist.Name)
		ar.Scan(&artist_id)
		al := db.QueryRow("INSERT INTO albums (title, artist_id) VALUES ($1, $2) ON CONFLICT (title, artist_id) DO UPDATE SET title=EXCLUDED.title returning id;", album.Title, artist_id)
		al.Scan(&album_id)
		ge := db.QueryRow("INSERT INTO genres (name) VALUES ($1) ON CONFLICT (name) DO UPDATE SET name=EXCLUDED.name returning name;", data.Genre())
		ge.Scan(&genre)
		_, err = db.Exec("INSERT INTO songs (title, track, comment, album_id, artist_id, genre) VALUES ($1,$2,$3,$4,$5,$6) ON CONFLICT (title, artist_id, album_id) DO UPDATE SET title=EXCLUDED.title returning id;", song.Title, song.Track, song.Comment, album_id, artist_id, genre)
		if err != nil {
			log.Fatal(err)
		}
		//
		tempFile, err := ioutil.TempFile("temp", "song-*.mp3")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)
		// return that we have successfully uploaded our file!
		fmt.Fprintf(w, "Successfully Uploaded File\n")
	}).Methods("POST")
}
