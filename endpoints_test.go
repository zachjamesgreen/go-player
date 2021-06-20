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

	expected := "[{\"id\":75,\"name\":\"A Place To Bury Strangers\"},{\"id\":87,\"name\":\"Broods\"},{\"id\":85,\"name\":\"BANKS\"},{\"id\":1,\"name\":\"Flume\"},{\"id\":21,\"name\":\"Marian Hill\"}]\n"
	ctype := res.Header["Content-Type"][0]
	Equal(t, ctype, "application/json")
	Equal(t, string(body), expected)

}

func TestGetArtist(t *testing.T) {
	db.Start()
	res, err := http.Get("http://localhost:8081/artists/87")
	chk(t,err)

	defer res.Body.Close()
	body,err := io.ReadAll(res.Body)
	chk(t,err)
	expected := "{\"id\":87,\"name\":\"Broods\"}\n"
	ctype := res.Header["Content-Type"][0]
	Equal(t, ctype, "application/json")
	Equal(t, string(body), expected)

}
