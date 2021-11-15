package models

import (
	"fmt"
	"testing"

	. "github.com/stretchr/testify/assert"
	"github.com/zachjamesgreen/go-player/database"
)

func TestStringSong(t *testing.T) {
	song := Song{ID: 1, Title: "test", Track: 1, Comment: "test", AlbumId: 1, ArtistId: 1, Path: "/test/path", Genre: "test Genre"}
	expected := "Song<ID: 1\n Title: test\n Track: 1\n Comment: test\n AlbumId: 1\n ArtistId: 1\n Path: /test/path\n Genre: test Genre>\n"
	Equal(t, expected, song.String())
}

func TestCreateSong(t *testing.T) {
	database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := Artist{Name: "test"}
	album := Album{Title: "test", Artist: &artist}
	song := Song{Title: "test", Album: &album, Artist: &artist}
	song.Create()
	NotEqual(t, 0, song.ID)
	NotEqual(t, 0, song.AlbumId)
	NotEqual(t, 0, song.ArtistId)
}

func TestFirstOrCreateSong(t *testing.T) {
	database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := Artist{Name: "test"}
	album := Album{Title: "test", Artist: &artist}
	song := Song{Title: "test", Album: &album, Artist: &artist}
	song.FirstOrCreate()
	id := song.ID
	NotEqual(t, 0, song.ID)
	NotEqual(t, 0, song.AlbumId)
	NotEqual(t, 0, song.ArtistId)
	song = Song{Title: "test", Album: &album, Artist: &artist}
	song.FirstOrCreate()
	Equal(t, id, song.ID)
}

func TestSaveSong(t *testing.T) {
	database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := Artist{Name: "test"}
	album := Album{Title: "test", Artist: &artist}
	song := Song{Title: "test", Album: &album, Artist: &artist}
	song.FirstOrCreate()
	song.Title = "test2"
	err := song.Save()
	NoError(t, err)
	Equal(t, "test2", song.Title)
}

func TestDeleteSong(t *testing.T) {
	db := database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := Artist{Name: "test"}
	album := Album{Title: "test", Artist: &artist}
	song := Song{Title: "test", Album: &album, Artist: &artist}
	song.FirstOrCreate()
	id := song.ID
	err := song.Delete()
	NoError(t, err)
	var expected Song
	result := db.Find(&expected, id)
	Equal(t, int64(0), result.RowsAffected)
	Equal(t, expected, Song{})
}

func TestGetSongs(t *testing.T) {
	database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := Artist{Name: "test"}
	album := Album{Title: "test", Artist: &artist}
	var songs []Song
	for i := 0; i < 10; i++ {
		title := fmt.Sprintf("test%d", i)
		song := Song{Title: title, Album: &album, Artist: &artist}
		song.FirstOrCreate()
		songs = append(songs, song)
	}

	expected, err := GetSongs()
	NoError(t, err)
	Equal(t, 10, len(expected))
	for i := 0; i < 10; i++ {
		Equal(t, songs[i].ID, expected[i].ID)
		Equal(t, songs[i].Title, expected[i].Title)
	}
}

func TestGetSong(t *testing.T) {
	database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := Artist{Name: "test"}
	album := Album{Title: "test", Artist: &artist}
	song := Song{Title: "test", Album: &album, Artist: &artist}
	song.FirstOrCreate()
	id := song.ID

	expected, err := GetSong(id)
	NoError(t, err)
	Equal(t, song.ID, expected.ID)
	Equal(t, song.Title, expected.Title)
}

func TestGetLikedSongs(t *testing.T) {
	database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := Artist{Name: "test"}
	album := Album{Title: "test", Artist: &artist}

	for i := 0; i < 10; i++ {
		title := fmt.Sprintf("test%d", i)
		liked := i%2 == 0
		song := Song{Title: title, Album: &album, Artist: &artist, Liked: liked}
		song.FirstOrCreate()
	}

	expected, err := GetLikedSongs()
	NoError(t, err)
	Equal(t, 5, len(expected))
	for _, song := range expected {
		Equal(t, true, song.Liked)
	}
}

func TestAddLike(t *testing.T) {
	database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := Artist{Name: "test"}
	album := Album{Title: "test", Artist: &artist}
	song := Song{Title: "test", Album: &album, Artist: &artist, Liked: false}
	song.Create()
	id := song.ID
	song.AddLike()
	expected, err := GetSong(id)
	NoError(t, err)
	Equal(t, true, expected.Liked)
}

func TestRemoveLike(t *testing.T) {
	database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := Artist{Name: "test"}
	album := Album{Title: "test", Artist: &artist}
	song := Song{Title: "test", Album: &album, Artist: &artist, Liked: true}
	song.Create()
	id := song.ID
	song.RemoveLike()
	expected, err := GetSong(id)
	NoError(t, err)
	Equal(t, false, expected.Liked)
}
