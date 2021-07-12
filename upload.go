package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	db "music/database"
	"music/models"

	"github.com/dhowden/tag"
	"gopkg.in/vansante/go-ffprobe.v2"
)

func UploadHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20)
	if files, ok := req.MultipartForm.File["song"]; ok {
		for _, fileHeader := range files {
			// if fileHeader.Header["Content-Type"][0] != "audio/mpeg" {
			// 	http.Error(w, "Can only handle audio/mpeg", http.StatusInternalServerError)
			// 	return
			// }
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

	artist_id, err := artist.Create()
	check(err)
	album_id, err := album.Create(artist_id)
	check(err)
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
	song.Duration = getDuration(full)
	err = ioutil.WriteFile(full, fileBytes, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		song.Create(artist_id, album_id, genre, full)
		// createSong(song, artist_id, album_id, genre, full)
	}

}

func getDuration(path string) uint64 {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	fileReader, err := os.Open(path)
	if err != nil {
		log.Panicf("Error opening test file: %v", err)
	}

	data, err := ffprobe.ProbeReader(ctx, fileReader)
	if err != nil {
		log.Panicf("Error getting data: %v", err)
	}
	return data.Streams[0].DurationTs
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
	// ext, err := mimeTypeToExt(data.Picture())
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	path := fmt.Sprintf("files/%s/%s/%s.%s", artist.Name, album.Title, "image", ext)
	// 	err = ioutil.WriteFile(path, data.Picture().Data, 0644)
	// 	if err != nil {
	// 		log.Println(err)
	// 		fmt.Println(err)
	// 	} else {
	// 		album.Image = true
	// 	}
	// }

	return artist, album, song, genreName
}

// func mimeTypeToExt(picture *tag.Picture) (mt string, err error) {
// 	err = nil
// 	if picture != nil {
// 		if picture.Ext != "" {
// 			mt = picture.Ext
// 		} else if picture.MIMEType == "image/jpg" {
// 			mt = "jpg"
// 		} else if picture.MIMEType == "image/jpeg" {
// 			mt = "jpg"
// 		}
// 	} else {
// 		err = fmt.Errorf("No picture")
// 	}
// 	return
// }

// func createArtist(artist models.Artist) string {
// 	var artist_id string
// 	ar := db.DB.QueryRow("INSERT INTO artists (name) VALUES ($1) ON CONFLICT (name) DO UPDATE SET name=EXCLUDED.name returning id;", artist.Name)
// 	err := ar.Scan(&artist_id)
// 	check(err)
// 	return artist_id
// }

// func createAlbum(album models.Album, artist_id string) string {
// 	var album_id string
// 	db.DB.QueryRow("INSERT INTO albums (title, artist_id, image) VALUES ($1, $2, $3) ON CONFLICT (title, artist_id) DO UPDATE SET title=EXCLUDED.title returning id;", album.Title, artist_id, album.Image).Scan(&album_id)
// 	return album_id
// }

func createGenre(genre models.Genre) string {
	var gen string
	db.DB.QueryRow("INSERT INTO genres (name) VALUES ($1) ON CONFLICT (name) DO UPDATE SET name=EXCLUDED.name returning name;", genre.Name).Scan(&gen)
	return gen
}

// func createSong(song models.Song, artist_id, album_id, genre, path string) {
// 	_, err := db.DB.Exec("INSERT INTO songs (title, track, comment, album_id, artist_id, genre, path, year, duration) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) ON CONFLICT (title, artist_id, album_id) DO UPDATE SET title=EXCLUDED.title returning id;", song.Title, song.Track, song.Comment, album_id, artist_id, genre, path, song.Year, song.duration)
// 	check(err)
// }
