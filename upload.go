package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"github.com/zachjamesgreen/go-player/models"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/dhowden/tag"
)

type SpotifyAlbumInfo struct {
	AlbumName  string `json:"album_name"`
	AlbumLink  string `json:"album_link"`
	ArtistName string `json:"artist_name"`
	ID         string `json:"id"`
	Images     string `json:"images"`
}

type Token struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   int       `json:"expires_in"`
	Expires     time.Time `json:"expires"`
}

var token Token

func UploadHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Panicf("Error parsing multipart form: %v", err)
	}
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
	artist, album, song, _ := getTagData(file)

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
		song.Upsert()
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
	log.Println("Got Token")
}

func getSpotifyAlbumArt(album *models.Album, artist models.Artist) {
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
		albumName := sAlbum.(map[string]interface{})["name"].(string)
		artistName := sAlbum.(map[string]interface{})["artists"].([]interface{})[0].(map[string]interface{})["name"].(string)
		if artistName == artist.Name && albumName == album.Title {
			album.SpotifyId = sAlbum.(map[string]interface{})["id"].(string)
			album.SpotifyLink = sAlbum.(map[string]interface{})["external_urls"].(map[string]interface{})["spotify"].(string)
			album.Images, err = json.Marshal(sAlbum.(map[string]interface{})["images"])
			check(err)
			return
		}
	}
}

func getSpotifyArtistArt(artist *models.Artist) {
	checkToken()

	parm := url.Values{}
	parm.Add("q", artist.Name)
	parm.Add("type", "artist")
	parm.Add("limit", "1")
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
	artist.SpotifyId = body.(map[string]interface{})["artists"].(map[string]interface{})["items"].([]interface{})[0].(map[string]interface{})["id"].(string)
	artist.Images, err = json.Marshal(body.(map[string]interface{})["artists"].(map[string]interface{})["items"].([]interface{})[0].(map[string]interface{})["images"])
	check(err)
	artist.Save()
}

func getTagData(file multipart.File) (artist models.Artist, album models.Album, song models.Song, genreName string) {
	data, err := tag.ReadFrom(file)
	check(err)

	artist = models.Artist{Name: data.Artist()}
	artist.FirstOrCreate()
	if artist.SpotifyId == "" {
		getSpotifyArtistArt(&artist)
	}

	album = models.Album{Title: data.Album(), Artist: &artist}
	album.Upsert()
	if album.SpotifyId == "" {
		getSpotifyAlbumArt(&album, artist)
		if err != nil {
			log.Panicf("Error marshalling images: %v", err)
		}
		album.Save()
	}

	track, _ := data.Track()
	genreName = data.Genre()
	song = models.Song{
		Title:   data.Title(),
		Track:   track,
		Comment: data.Comment(),
		Genre:   genreName,
		Artist:  &artist,
		Album:   &album,
		Year:    data.Year(),
	}
	return
}
