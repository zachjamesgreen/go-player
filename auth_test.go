package main

import (
	"net/http"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	// Test with out user or password
	res, err := http.Get("http://localhost:8081/artists")
	chk(t, err)
	Equal(t, res.StatusCode, 403)

	// Test with user but no password
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8081/artists", nil)
	chk(t, err)

	req.SetBasicAuth("zach", "")
	res, err = client.Do(req)
	chk(t, err)
	Equal(t, res.StatusCode, 401)

	// Test with password but no user

	req.SetBasicAuth("", "1234")
	res, err = client.Do(req)
	chk(t, err)
	Equal(t, res.StatusCode, 401)

	// Test with both user and password

	req.SetBasicAuth("zach", "1234")
	res, err = client.Do(req)
	chk(t, err)
	Equal(t, res.StatusCode, 200)
}
