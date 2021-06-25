package main

import (
	"io"
	"io/ioutil"
	"log"
	db "music/database"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func chk(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func DBSetUp() {
	path := filepath.Join("tables.sql")

	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		log.Fatal(ioErr)
	}
	sql := string(c)
	db.DB.Exec("TRUNCATE TABLE songs")
	_, err := db.DB.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	log.Println("Starting Up")
	db.Start()
	DBSetUp()

	exitVal := m.Run()
	log.Println("Done!")

	os.Exit(exitVal)
}

func TestGetArtists(t *testing.T) {

	res, err := http.Get("http://localhost:8081/artists")
	chk(t, err)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	chk(t, err)

	expected := "[{\"id\":1,\"name\":\"Broods\"},{\"id\":2,\"name\":\"BANKS\"},{\"id\":3,\"name\":\"Flume\"},{\"id\":4,\"name\":\"Sleigh Bells\"},{\"id\":5,\"name\":\"Explosions In The Sky\"},{\"id\":6,\"name\":\"Three\"}]\n"
	ctype := res.Header["Content-Type"][0]
	Equal(t, ctype, "application/json")
	Equal(t, string(body), expected)

}

func TestGetArtist(t *testing.T) {
	db.Start()
	res, err := http.Get("http://localhost:8081/artists/2")
	chk(t, err)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	chk(t, err)
	expected := "{\"id\":2,\"name\":\"BANKS\"}\n"
	ctype := res.Header["Content-Type"][0]
	Equal(t, ctype, "application/json")
	Equal(t, string(body), expected)
}

func TestGetArtistSongs(t *testing.T) {
	db.Start()
	res, err := http.Get("http://localhost:8081/artists/2/songs")
	chk(t, err)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	chk(t, err)
	expected := "[{\"id\":2,\"title\":\"Gemini Feed\",\"track\":1,\"comment\":\"\",\"year\":2016,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":2,\"album_id\":2,\"artist\":\"BANKS\",\"album\":\"The Altar\",\"path\":\"files/BANKS/The Altar/01 Gemini Feed.mp3\",\"genre\":{\"name\":\"Alternative/Indie\"}},{\"id\":6,\"title\":\"Fuck With Myself\",\"track\":2,\"comment\":\"\",\"year\":2016,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":2,\"album_id\":2,\"artist\":\"BANKS\",\"album\":\"The Altar\",\"path\":\"files/BANKS/The Altar/02 Fuck With Myself.mp3\",\"genre\":{\"name\":\"Alternative/Indie\"}}]\n"
	ctype := res.Header["Content-Type"][0]
	Equal(t, ctype, "application/json")
	Equal(t, string(body), expected)
}

func TestGetArtistAlbums(t *testing.T) {
	db.Start()
	res, err := http.Get("http://localhost:8081/artists/2/albums")
	chk(t, err)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	chk(t, err)
	expected := "[{\"id\":2,\"title\":\"The Altar\",\"artist_id\":2}]\n"
	ctype := res.Header["Content-Type"][0]
	Equal(t, ctype, "application/json")
	Equal(t, string(body), expected)
}

func TestGetAlbums(t *testing.T) {
	db.Start()
	res, err := http.Get("http://localhost:8081/albums")
	chk(t, err)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	chk(t, err)
	expected := "[{\"id\":1,\"title\":\"Conscious\",\"artist_id\":1},{\"id\":2,\"title\":\"The Altar\",\"artist_id\":2},{\"id\":3,\"title\":\"Skin\",\"artist_id\":3},{\"id\":4,\"title\":\"Jessica Rabbit\",\"artist_id\":4},{\"id\":5,\"title\":\"The Wilderness\",\"artist_id\":5},{\"id\":6,\"title\":\"Phantogram\",\"artist_id\":6}]\n"
	ctype := res.Header["Content-Type"][0]
	Equal(t, ctype, "application/json")
	Equal(t, string(body), expected)
}

func TestGetAlbum(t *testing.T) {
	db.Start()
	res, err := http.Get("http://localhost:8081/albums/2")
	chk(t, err)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	chk(t, err)
	expected := "{\"id\":2,\"title\":\"The Altar\",\"artist_id\":2}\n"
	ctype := res.Header["Content-Type"][0]
	Equal(t, ctype, "application/json")
	Equal(t, string(body), expected)
}

func TestGetAlbumSongs(t *testing.T) {
	db.Start()
	res, err := http.Get("http://localhost:8081/albums/2/songs")
	chk(t, err)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	chk(t, err)
	expected := "[{\"id\":2,\"title\":\"Gemini Feed\",\"track\":1,\"comment\":\"\",\"year\":2016,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":2,\"album_id\":2,\"artist\":\"BANKS\",\"album\":\"The Altar\",\"path\":\"files/BANKS/The Altar/01 Gemini Feed.mp3\",\"genre\":{\"name\":\"Alternative/Indie\"}},{\"id\":6,\"title\":\"Fuck With Myself\",\"track\":2,\"comment\":\"\",\"year\":2016,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":2,\"album_id\":2,\"artist\":\"BANKS\",\"album\":\"The Altar\",\"path\":\"files/BANKS/The Altar/02 Fuck With Myself.mp3\",\"genre\":{\"name\":\"Alternative/Indie\"}}]\n"
	ctype := res.Header["Content-Type"][0]
	Equal(t, ctype, "application/json")
	Equal(t, string(body), expected)
}

func TestGetSongs(t *testing.T) {
	db.Start()
	res, err := http.Get("http://localhost:8081/songs")
	chk(t, err)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	chk(t, err)
	expected := "[{\"id\":1,\"title\":\"Free\",\"track\":1,\"comment\":\"\",\"year\":2016,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":1,\"album_id\":1,\"artist\":\"Broods\",\"album\":\"Conscious\",\"path\":\"files/Broods/Conscious/01 Free.mp3\",\"genre\":{\"name\":\"Alternative/Indie\"}},{\"id\":2,\"title\":\"Gemini Feed\",\"track\":1,\"comment\":\"\",\"year\":2016,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":2,\"album_id\":2,\"artist\":\"BANKS\",\"album\":\"The Altar\",\"path\":\"files/BANKS/The Altar/01 Gemini Feed.mp3\",\"genre\":{\"name\":\"Alternative/Indie\"}},{\"id\":3,\"title\":\"Helix\",\"track\":1,\"comment\":\"\",\"year\":0,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":3,\"album_id\":3,\"artist\":\"Flume\",\"album\":\"Skin\",\"path\":\"files/Flume/Skin/01 Helix.mp3\",\"genre\":{\"name\":\"Dance/Electronic\"}},{\"id\":4,\"title\":\"It's Just Us Now\",\"track\":1,\"comment\":\"\",\"year\":0,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":4,\"album_id\":4,\"artist\":\"Sleigh Bells\",\"album\":\"Jessica Rabbit\",\"path\":\"files/Sleigh Bells/Jessica Rabbit/01 It's Just Us Now.mp3\",\"genre\":{\"name\":\"Alternative/Indie\"}},{\"id\":5,\"title\":\"Wilderness\",\"track\":1,\"comment\":\"\",\"year\":2016,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":5,\"album_id\":5,\"artist\":\"Explosions In The Sky\",\"album\":\"The Wilderness\",\"path\":\"files/Explosions In The Sky/The Wilderness/01 Wilderness.mp3\",\"genre\":{\"name\":\"Alternative/Indie\"}},{\"id\":6,\"title\":\"Fuck With Myself\",\"track\":2,\"comment\":\"\",\"year\":2016,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":2,\"album_id\":2,\"artist\":\"BANKS\",\"album\":\"The Altar\",\"path\":\"files/BANKS/The Altar/02 Fuck With Myself.mp3\",\"genre\":{\"name\":\"Alternative/Indie\"}},{\"id\":7,\"title\":\"Never Be Like You (feat. Kai)\",\"track\":2,\"comment\":\"\",\"year\":0,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":3,\"album_id\":3,\"artist\":\"Flume\",\"album\":\"Skin\",\"path\":\"files/Flume/Skin/02 Never Be Like You (feat_ Kai).mp3\",\"genre\":{\"name\":\"Dance/Electronic\"}},{\"id\":8,\"title\":\"The Ecstatics\",\"track\":2,\"comment\":\"\",\"year\":0,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":5,\"album_id\":5,\"artist\":\"Explosions In The Sky\",\"album\":\"The Wilderness\",\"path\":\"files/Explosions In The Sky/The Wilderness/02 The Ecstatics.mp3\",\"genre\":{\"name\":\"Alternative/Indie\"}},{\"id\":9,\"title\":\"Torn Clean\",\"track\":2,\"comment\":\"\",\"year\":2016,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":4,\"album_id\":4,\"artist\":\"Sleigh Bells\",\"album\":\"Jessica Rabbit\",\"path\":\"files/Sleigh Bells/Jessica Rabbit/02 Torn Clean.mp3\",\"genre\":{\"name\":\"Alternative/Indie\"}},{\"id\":10,\"title\":\"We Had Everything\",\"track\":2,\"comment\":\"\",\"year\":2016,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":1,\"album_id\":1,\"artist\":\"Broods\",\"album\":\"Conscious\",\"path\":\"files/Broods/Conscious/02 We Had Everything.mp3\",\"genre\":{\"name\":\"Alternative/Indie\"}},{\"id\":11,\"title\":\"Funeral Pyre\",\"track\":1,\"comment\":\"\",\"year\":2016,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":6,\"album_id\":6,\"artist\":\"Three\",\"album\":\"Phantogram\",\"path\":\"files/Phantogram/Three/01 Funeral Pyre.mp3\",\"genre\":{\"name\":\"Alternative/Indie\"}},{\"id\":12,\"title\":\"Same Old Blues\",\"track\":2,\"comment\":\"\",\"year\":2016,\"last_played\":{\"Time\":\"0001-01-01T00:00:00Z\",\"Valid\":false},\"artist_id\":6,\"album_id\":6,\"artist\":\"Three\",\"album\":\"Phantogram\",\"path\":\"files/Phantogram/Three/02 Same Old Blues.mp3\",\"genre\":{\"name\":\"Alternative/Indie\"}}]\n"
	ctype := res.Header["Content-Type"][0]
	Equal(t, ctype, "application/json")
	Equal(t, string(body), expected)
}
