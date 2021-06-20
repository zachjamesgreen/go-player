package main

import (
	"io"
	db "music/database"
	"net/http"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func chk(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetArtists(t *testing.T) {
	db.Start()
	res, err := http.Get("http://localhost:8081/artists")
	chk(t,err)
	
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	chk(t,err)

	expected := "[{\"id\":2,\"name\":\"BANKS\"},{\"id\":3,\"name\":\"Flume\"},{\"id\":5,\"name\":\"Explosions In The Sky\"},{\"id\":4,\"name\":\"Sleigh Bells\"},{\"id\":1,\"name\":\"Broods\"}]\n"
	ctype := res.Header["Content-Type"][0]
	Equal(t, ctype, "application/json")
	Equal(t, string(body), expected)

}

func TestGetArtist(t *testing.T) {
	db.Start()
	res, err := http.Get("http://localhost:8081/artists/2")
	chk(t,err)

	defer res.Body.Close()
	body,err := io.ReadAll(res.Body)
	chk(t,err)
	expected := "{\"id\":2,\"name\":\"BANKS\"}\n"
	ctype := res.Header["Content-Type"][0]
	Equal(t, ctype, "application/json")
	Equal(t, string(body), expected)
}

func TestGetArtistSongs(t *testing.T) {
	db.Start()
	res, err := http.Get("http://localhost:8081/artists/2/songs")
	chk(t,err)

	defer res.Body.Close()
	body,err := io.ReadAll(res.Body)
	chk(t,err)
	expected := "[{\"id\":6,\"title\":\"Fuck With Myself\",\"track\":2,\"comment\":\"\",\"artist_id\":2,\"album_id\":2,\"path\":\"files/BANKS/The Altar\",\"genre\":{\"name\":\"Alternative/Indie\"}},{\"id\":2,\"title\":\"Gemini Feed\",\"track\":1,\"comment\":\"\",\"artist_id\":2,\"album_id\":2,\"path\":\"files/BANKS/The Altar\",\"genre\":{\"name\":\"Alternative/Indie\"}}]\n"
	ctype := res.Header["Content-Type"][0]
	Equal(t, ctype, "application/json")
	Equal(t, string(body), expected)
}

func TestGetArtistAlbums(t *testing.T) {
	db.Start()
	res, err := http.Get("http://localhost:8081/artists/2/albums")
	chk(t,err)

	defer res.Body.Close()
	body,err := io.ReadAll(res.Body)
	chk(t,err)
	expected := "[{\"id\":2,\"title\":\"The Altar\",\"artist_id\":2}]\n"
	ctype := res.Header["Content-Type"][0]
	Equal(t, ctype, "application/json")
	Equal(t, string(body), expected)
}