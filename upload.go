package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	db "music/database"
	"music/models"

	"github.com/dhowden/tag"
	"gopkg.in/vansante/go-ffprobe.v2"
)

type SpotifyAlbumInfo struct {
	AlbumName  string `json:"album_name"`
	AlbumLink  string `json:"album_link"`
	ArtistName string `json:"artist_name"`
	ID         string `json:"id"`
	Images     string `json:"images"`
}

// type SpotifyImages struct {
// 	Height int    `json:"height"`
// 	Width  int    `json:"width"`
// 	URL    string `json:"url"`
// }

type Token struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   int       `json:"expires_in"`
	Expires     time.Time `json:"expires"`
}

var token Token

func UploadHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Token", token)
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
	album_id, err := album.Upsert(artist_id)
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
	err = ioutil.WriteFile(full, fileBytes, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		// song.Duration = getDuration(full)
		song.Create(artist_id, album_id, genre, full)
	}
}

func checkToken() {
	if token.Expires.After(time.Now()) {
		log.Println("Token Good")
		return
	}
	log.Println("Getting Token")

	parm := url.Values{}
	parm.Add("grant_type", "client_credentials")
	req, err := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(parm.Encode()))
	if err != nil {
		log.Panicf("Error creating request: %v", err)
	}

	req.SetBasicAuth(os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panicf("Error sending request: %v", err)
	}
	if res.StatusCode != 200 {
		log.Panicf("Error: %v", res.Body)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicf("Error reading response body: %v", err)
	}

	err = json.Unmarshal(resBody, &token)
	if err != nil {
		log.Panicf("Error unmarshalling response: %v", err)
	}
	token.Expires = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	log.Println("Got Token", token)
}

func getSpotifyAlbumArt(album models.Album, artist models.Artist) SpotifyAlbumInfo {
	var albumInfo SpotifyAlbumInfo
	checkToken()

	parm := url.Values{}
	parm.Add("q", album.Title)
	parm.Add("type", "album")
	req, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/search?"+parm.Encode(), http.NoBody)
	if err != nil {
		log.Panicf("Error creating request: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panicf("Error sending request: %v", err)
	}
	if res.StatusCode != 200 {
		log.Panicf("Error: %v", res.Body)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicf("Error reading response body: %v", err)
	}
	var body interface{}
	err = json.Unmarshal(resBody, &body)
	if err != nil {
		log.Panicf("Error unmarshalling response: %v", err)
	}
	for _, sAlbum := range body.(map[string]interface{})["albums"].(map[string]interface{})["items"].([]interface{}) {
		albumInfo.AlbumName = sAlbum.(map[string]interface{})["name"].(string)
		albumInfo.AlbumLink = sAlbum.(map[string]interface{})["external_urls"].(map[string]interface{})["spotify"].(string)
		albumInfo.ID = sAlbum.(map[string]interface{})["id"].(string)
		str, err := json.Marshal(sAlbum.(map[string]interface{})["images"])
		check(err)
		albumInfo.Images = string(str)
		albumInfo.ArtistName = sAlbum.(map[string]interface{})["artists"].([]interface{})[0].(map[string]interface{})["name"].(string)
		if albumInfo.ArtistName == artist.Name && albumInfo.AlbumName == album.Title {
			return albumInfo
		}
	}
	return albumInfo
}

func getSpotifyArtistArt(artist models.Artist) {}

func getDuration(path string) uint64 {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	fmt.Println(path)

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
	albumInfo := getSpotifyAlbumArt(album, artist)
	album.SpotifyId = albumInfo.ID
	album.SpotifyLink = albumInfo.AlbumLink
	album.Images = albumInfo.Images
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
