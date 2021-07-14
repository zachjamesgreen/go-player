package models

import (
	"database/sql"
	"fmt"
	"log"
	db "music/database"
	"time"
)

type LikedSong struct {
	ID        int       `json:"id"`
	Song      Song      `json:"song"`
	DateAdded time.Time `json:"date_added"`
}

func (ls *LikedSong) All() (lss []LikedSong, err error) {
	ls.Song = Song{}
	sqlStatment := `
	SELECT s.id, s.title, s.track, 
				s.comment, s.year, s.last_played, 
				s.path, s.genre, s.album_id, 
				s.artist_id, s.created_at, s.duration, 
				al.title as album_title, ar.name, 
				ls.date_added AS added_to_ls
	FROM songs AS s
	JOIN liked_songs AS ls ON ls.song_id = s.id
	JOIN albums AS al ON s.album_id = al.id
	JOIN artists AS ar ON s.artist_id = ar.id ORDER BY s.id;`
	rows, err := db.DB.Query(sqlStatment)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows")
		} else {
			panic(err)
		}
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&ls.Song.Id, &ls.Song.Title, &ls.Song.Track,
			&ls.Song.Comment, &ls.Song.Year, &ls.Song.LastPlayed,
			&ls.Song.Path, &ls.Song.Genre.Name, &ls.Song.AlbumId,
			&ls.Song.ArtistId, &ls.Song.CreatedAt, &ls.Song.Duration,
			&ls.Song.Album, &ls.Song.Artist, &ls.DateAdded)
		if err != nil {
			log.Fatal(err)
		}
		lss = append(lss, *ls)
	}
	return
}

func (ls *LikedSong) Add(id int) (err error) {
	sqlStatment := `
	INSERT INTO liked_songs (song_id, date_added)
	VALUES ($1, $2);`
	_, err = db.DB.Exec(sqlStatment, id, time.Now())
	return
}

func (ls *LikedSong) Remove(id int) (err error) {
	sqlStatment := `
	DELETE FROM liked_songs WHERE song_id = $1;`
	_, err = db.DB.Exec(sqlStatment, id)
	return
}
