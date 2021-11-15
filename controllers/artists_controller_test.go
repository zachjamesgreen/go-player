package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	. "github.com/stretchr/testify/assert"
	"github.com/zachjamesgreen/go-player/database"
	"github.com/zachjamesgreen/go-player/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGetArtists(t *testing.T) {
	database.GetTestDB(false)
	defer database.CleanTestDB()
	artists := []models.Artist{
		{
			Name: "Test1",
		},
		{
			Name: "Test2",
		},
		{
			Name: "Test3",
		},
	}

	for _, artist := range artists {
		err := artist.FirstOrCreate()
		NoError(t, err)
	}

	req, _ := http.NewRequest("GET", "/artists", nil)
	res := httptest.NewRecorder()
	GetArtists(res, req)

	var dest []interface{}
	err := json.Unmarshal(res.Body.Bytes(), &dest)
	NoError(t, err)

	Equal(t, 200, res.Code)
	Equal(t, 3, len(dest))
}

func TestGetArtist(t *testing.T) {
	database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := models.Artist{
		Name: "Test1",
	}
	err := artist.FirstOrCreate()
	NoError(t, err)

	url := fmt.Sprintf("/artists/%d", artist.ID)
	req, _ := http.NewRequest("GET", url, nil)
	res := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/artists/{id}", GetArtist)
	router.ServeHTTP(res, req)

	var dest models.Artist
	err = json.Unmarshal(res.Body.Bytes(), &dest)
	NoError(t, err)

	Equal(t, 200, res.Code)
	Equal(t, artist.Name, dest.Name)
	Equal(t, artist.ID, dest.ID)
}

func TestGetArtistSongs(t *testing.T) {
	database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := models.Artist{
		Name: "Test1",
	}
	album := models.Album{
		Title:  "Test1",
		Artist: &artist,
	}
	err := artist.FirstOrCreate()
	NoError(t, err)
	err = album.Upsert()
	NoError(t, err)
	var songs []models.Song
	for i := 0; i < 3; i++ {
		song := models.Song{Title: fmt.Sprintf("Test%d", i), Album: &album, Artist: &artist}
		err = song.Upsert()
		NoError(t, err)
		songs = append(songs, song)
	}

	url := fmt.Sprintf("/artists/%d/songs", artist.ID)
	req, _ := http.NewRequest("GET", url, nil)
	res := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/artists/{id}/songs", GetArtistSongs)
	router.ServeHTTP(res, req)

	var dest []models.Song
	err = json.Unmarshal(res.Body.Bytes(), &dest)
	NoError(t, err)

	Equal(t, 200, res.Code)
	Equal(t, 3, len(dest))
	for i := 0; i < 3; i++ {
		Equal(t, songs[i].Title, dest[i].Title)
	}
}

func TestGetArtistAlbums(t *testing.T) {
	database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := models.Artist{
		Name: "Test1",
	}
	err := artist.FirstOrCreate()
	NoError(t, err)
	var albums []models.Album
	for i := 0; i < 3; i++ {
		album := models.Album{Title: fmt.Sprintf("Test%d", i), Artist: &artist}
		err = album.Upsert()
		NoError(t, err)
		albums = append(albums, album)
	}

	url := fmt.Sprintf("/artists/%d/albums", artist.ID)
	res := SendRequest("/artists/{id}/albums", url, "GET", GetArtistAlbums)

	var dest []models.Album
	err = json.Unmarshal(res.Body.Bytes(), &dest)
	NoError(t, err)

	Equal(t, 200, res.Code)
	Equal(t, 3, len(dest))
	for i := 0; i < 3; i++ {
		Equal(t, albums[i].Title, dest[i].Title)
	}
}

func TestDeleteArtist(t *testing.T) {
	db := database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := models.Artist{
		Name: "Test1",
	}
	err := artist.FirstOrCreate()
	NoError(t, err)

	url := fmt.Sprintf("/artists/%d", artist.ID)
	res := SendRequest("/artists/{id}", url, "DELETE", DeleteArtist)


	
	Equal(t, 200, res.Code)
	expected := models.Artist{}
	result := db.First(&expected, artist.ID)
	Equal(t, result.Error, gorm.ErrRecordNotFound)
	
	// TODO: test for 404 if artist does not exist
}

func TestDeleteArtistIfNotExists(t *testing.T) {
	database.GetTestDB(false)
	defer database.CleanTestDB()

	url := fmt.Sprintf("/artists/%d", 1000)
	res := SendRequest("/artists/{id}", url, "DELETE", DeleteArtist)

	Equal(t, 404, res.Code)
}
