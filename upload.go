package main

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/dhowden/tag"
)

func upload(file multipart.File, handler *multipart.FileHeader) {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(db)

	// Get tags
	artist, album, song, genreName := getTagData(file)

	artist_id := createArtist(artist)
	album_id := createAlbum(album, artist_id)
	genre := createGenre(genreName)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	path := fmt.Sprintf("temp/%s/%s", artist.Name, album.Title)
	err = os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println(err)
	}
	full := fmt.Sprintf("%s/%s", path, handler.Filename)
	err = ioutil.WriteFile(full, fileBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}

	createSong(song, artist_id, album_id, genre, path)
}

func getTagData(file multipart.File) (Artist, Album, Song, Genre) {
	data, err := tag.ReadFrom(file)
	check(err)

	var artist Artist = Artist{Name: data.Artist()}
	var album Album = Album{Title: data.Album(), ArtistId: 0}
	track, _ := data.Track()
	genreName := Genre{Name: data.Genre()}
	var song Song = Song{Title: data.Title(), Track: track, Comment: data.Comment(), Genre: genreName, ArtistId: 0, AlbumId: 0}
	return artist, album, song, genreName
}

func createArtist(artist Artist) string {
	var artist_id string
	ar := db.QueryRow("INSERT INTO artists (name) VALUES ($1) ON CONFLICT (name) DO UPDATE SET name=EXCLUDED.name returning id;", artist.Name)
	err := ar.Scan(&artist_id)
	check(err)
	return artist_id
}

func createAlbum(album Album, artist_id string) string {
	var album_id string
	db.QueryRow("INSERT INTO albums (title, artist_id) VALUES ($1, $2) ON CONFLICT (title, artist_id) DO UPDATE SET title=EXCLUDED.title returning id;", album.Title, artist_id).Scan(&album_id)
	return album_id
}

func createGenre(genre Genre) string {
	var gen string
	db.QueryRow("INSERT INTO genres (name) VALUES ($1) ON CONFLICT (name) DO UPDATE SET name=EXCLUDED.name returning name;", genre.Name).Scan(&gen)
	return gen
}

func createSong(song Song, artist_id, album_id, genre, path string) {
	_, err := db.Exec("INSERT INTO songs (title, track, comment, album_id, artist_id, genre, path) VALUES ($1,$2,$3,$4,$5,$6,$7) ON CONFLICT (title, artist_id, album_id) DO UPDATE SET title=EXCLUDED.title returning id;", song.Title, song.Track, song.Comment, album_id, artist_id, genre, path)
	check(err)
}
