package models

import (
	. "github.com/stretchr/testify/assert"
	"music/database"
	"testing"
)

func TestArtistString(t *testing.T) {
	expected := "Artist<ID: 1 | Name: Test>\n"
	artist := Artist{
		ID:   1,
		Name: "Test",
	}
	Equal(t, expected, artist.String())
}

func TestFirstOrCreateError(t *testing.T) {
	artist := Artist{}
	database.GetTestDB(false)
	defer database.CleanTestDB()

	err := artist.FirstOrCreate()
	Error(t, err)
}

func TestFirstOrCreate(t *testing.T) {
	db := database.GetTestDB(false)
	defer database.CleanTestDB()
	var artist Artist
	expected := Artist{
		Name: "Test",
	}

	err := expected.FirstOrCreate()
	NoError(t, err)
	result := db.First(&artist)
	NoError(t, result.Error)
	Equal(t, expected.Name, artist.Name)

	id := artist.ID

	err = artist.FirstOrCreate()
	NoError(t, err)
	result = db.First(&artist, artist.ID)
	NoError(t, result.Error)
	Equal(t, id, artist.ID)
}

func TestSave(t *testing.T) {
	db := database.GetTestDB(false)
	defer database.CleanTestDB()
	var artist Artist
	expected := Artist{
		ID:   2,
		Name: "Test",
	}

	err := expected.Save()
	NoError(t, err)
	result := db.First(&artist)
	NoError(t, result.Error)
	Equal(t, expected, artist)
	Equal(t, 2, artist.ID)

	// error if name is changes to empty string
	artist.Name = ""
	err = artist.Save()
	Error(t, err)
}

func TestDelete(t *testing.T) {
	artist := Artist{
		Name: "Test",
	}
	database.GetTestDB(false)
	defer database.CleanTestDB()

	err := artist.FirstOrCreate()
	NoError(t, err)

	err = artist.Delete()
	NoError(t, err)

	artist, err = GetArtistById(artist.ID)
	NoError(t, err)
	Equal(t, "", artist.Name)
	// test with albums and songs
}

func TestGetAllArtists(t *testing.T) {
	artists := []Artist{
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
	database.GetTestDB(false)
	defer database.CleanTestDB()

	for _, artist := range artists {
		err := artist.FirstOrCreate()
		NoError(t, err)
	}

	artists, err := GetAllArtists()
	NoError(t, err)
	Equal(t, 3, len(artists))
}

func TestGetArtistById(t *testing.T) {
	artist := Artist{
		Name: "Test",
	}
	database.GetTestDB(false)
	defer database.CleanTestDB()

	err := artist.FirstOrCreate()
	NoError(t, err)

	expected, err := GetArtistById(artist.ID)
	NoError(t, err)
	Equal(t, "Test", expected.Name)
}

func TestGetArtistAlbumsById(t *testing.T) {
	artist := Artist{
		Name: "Test",
	}

	albums := []Album{
		{
			Title: "Test1",
			Artist: &artist,
		},
		{
			Title: "Test2",
			Artist: &artist,
		},
		{
			Title: "Test3",
			Artist: &artist,
		},
	}

	database.GetTestDB(false)
	defer database.CleanTestDB()

	err := artist.FirstOrCreate()
	NoError(t, err)

	for _, album := range albums {
		album.Upsert()
		NoError(t, err)
	}

	artistAlbums, err := GetArtistAlbumsById(artist.ID)
	NoError(t, err)
	Equal(t, 3, len(artistAlbums))
}

func TestGetArtistSongsById(t *testing.T) {
	artist := Artist{
		Name: "Test",
	}
	album := Album{
		Title: "Test",
		Artist: &artist,
	}
	songs := []Song{
		{
			Title: "Test1",
			Album: &album,
			Artist: &artist,
		},
		{
			Title: "Test2",
			Album: &album,
			Artist: &artist,
		},
		{
			Title: "Test3",
			Album: &album,
			Artist: &artist,
		},
	}


	database.GetTestDB(false)
	defer database.CleanTestDB()

	err := artist.FirstOrCreate()
	NoError(t, err)
	err = album.Upsert()
	for _, song := range songs {
		song.Upsert()
		NoError(t, err)
	}

	result, err := GetArtistSongsById(artist.ID)
	NoError(t, err)
	Equal(t, 3, len(result))
}

