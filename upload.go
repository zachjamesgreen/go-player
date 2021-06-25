package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	db "music/database"
	"music/models"

	"github.com/dhowden/tag"
)

func UploadHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20)
	if files, ok := req.MultipartForm.File["song"]; ok {
		for _, fileHeader := range files {
			if fileHeader.Header["Content-Type"][0] != "audio/mpeg" {
				http.Error(w, "Can only handle audio/mpeg", http.StatusInternalServerError)
				return
			}
			fmt.Printf("Uploaded File: %+v\n", fileHeader.Filename)
			fmt.Printf("File Size: %+v\n", fileHeader.Size)
			fmt.Printf("MIME Header: %+v\n", fileHeader.Header["Content-Type"][0])
			upload(fileHeader)
		}

		// return that we have successfully uploaded our file!
		fmt.Fprintf(w, "Successfully Uploaded File\n")
	} else {
		http.Error(w, "Unprocessable Entity", http.StatusUnprocessableEntity)
		return
	}
}

func upload(fileHeader *multipart.FileHeader) {
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println(err)
	}

	// Get tags
	artist, album, song, genreName := getTagData(file)

	artist_id := createArtist(artist)
	album_id := createAlbum(album, artist_id)
	genre := createGenre(genreName)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	path := fmt.Sprintf("files/%s/%s", artist.Name, album.Title)
	err = os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println(err)
	}
	full := fmt.Sprintf("%s/%s", path, fileHeader.Filename)
	err = ioutil.WriteFile(full, fileBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}

	createSong(song, artist_id, album_id, genre, full)
}

func getTagData(file multipart.File) (models.Artist, models.Album, models.Song, models.Genre) {
	data, err := tag.ReadFrom(file)
	check(err)

	var artist = models.Artist{Name: data.Artist()}
	var album = models.Album{Title: data.Album(), ArtistId: 0}
	track, _ := data.Track()
	genreName := models.Genre{Name: data.Genre()}
	var song = models.Song{Title: data.Title(), Track: track, Comment: data.Comment(), Genre: genreName, ArtistId: 0, AlbumId: 0, Year: data.Year()}

	// Save Image
	log.Print(data.Picture().Ext == "")
	path := fmt.Sprintf("files/%s/%s/%s.%s", artist.Name, album.Title, "image", mimeTypeToExt(data.Picture()))
	err = ioutil.WriteFile(path, data.Picture().Data, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		album.Image = true
	}

	return artist, album, song, genreName
}

func mimeTypeToExt(picture *tag.Picture) string {
	var mt string
	if picture.Ext != "" {
		mt = picture.Ext
	} else if picture.MIMEType == "image/jpg" {
		mt = "jpg"
	} else if picture.MIMEType == "image/jpeg" {
		mt = "jpg"
	}
	return mt
}

func createArtist(artist models.Artist) string {
	var artist_id string
	ar := db.DB.QueryRow("INSERT INTO artists (name) VALUES ($1) ON CONFLICT (name) DO UPDATE SET name=EXCLUDED.name returning id;", artist.Name)
	err := ar.Scan(&artist_id)
	check(err)
	return artist_id
}

func createAlbum(album models.Album, artist_id string) string {
	var album_id string
	db.DB.QueryRow("INSERT INTO albums (title, artist_id, image) VALUES ($1, $2, $3) ON CONFLICT (title, artist_id) DO UPDATE SET title=EXCLUDED.title returning id;", album.Title, artist_id, album.Image).Scan(&album_id)
	return album_id
}

func createGenre(genre models.Genre) string {
	var gen string
	db.DB.QueryRow("INSERT INTO genres (name) VALUES ($1) ON CONFLICT (name) DO UPDATE SET name=EXCLUDED.name returning name;", genre.Name).Scan(&gen)
	return gen
}

func createSong(song models.Song, artist_id, album_id, genre, path string) {
	_, err := db.DB.Exec("INSERT INTO songs (title, track, comment, album_id, artist_id, genre, path, year) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) ON CONFLICT (title, artist_id, album_id) DO UPDATE SET title=EXCLUDED.title returning id;", song.Title, song.Track, song.Comment, album_id, artist_id, genre, path, song.Year)
	check(err)
}
