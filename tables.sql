TRUNCATE TABLE artists CASCADE;
TRUNCATE TABLE genres CASCADE;

ALTER SEQUENCE albums_id_seq RESTART WITH 1;
ALTER SEQUENCE artists_id_seq RESTART WITH 1;
ALTER SEQUENCE songs_id_seq RESTART WITH 1;

-- create table artists (
--   id serial primary key,
--   name varchar UNIQUE
--   );

-- create table albums (
--   id serial primary key,
--   title varchar,
--   artist_id integer references artists(id),
--   UNIQUE (title, artist_id)
--   );

-- create table genres (
--   name varchar primary key UNIQUE
--   );

-- create table songs (
--   id serial primary key,
--   title varchar,
--   track int,
--   comment text,
--   album_id integer references albums(id),
--   artist_id integer references artists(id),
--   genre varchar references genres(name),
--   path varchar,
--   UNIQUE (title, artist_id, album_id)
--   );

-- create table users (
--   id serial primary key,
--   username varchar NOT NULL,
--   password varchar NOT NULL,
--   created_at timestamp,
--   updated_at timestamp,
--   UNIQUE (username)
-- );

-- alter table songs add column last_played timestamp;
-- alter table songs add column year int;

INSERT INTO artists (id,name) VALUES
(1, 'Broods'),
(2, 'BANKS'),
(3, 'Flume'),
(4, 'Sleigh Bells'),
(5, 'Explosions In The Sky'),
(6, 'Three');

INSERT INTO albums (id,title,artist_id) VALUES
(1, 'Conscious', 1),
(2, 'The Altar', 2),
(3, 'Skin', 3),
(4, 'Jessica Rabbit', 4),
(5, 'The Wilderness', 5),
(6, 'Phantogram', 6);

INSERT INTO genres (name) VALUES
('Dance/Electronic'),
('Alternative/Indie');

INSERT INTO songs (id,title,track,comment,album_id,artist_id,genre,path,year) VALUES
(1,'Free',1,'',1,1,'Alternative/Indie','files/Broods/Conscious/01 Free.mp3',2016),
(2,'Gemini Feed',1,'',2,2,'Alternative/Indie','files/BANKS/The Altar/01 Gemini Feed.mp3',2016),
(3,'Helix',1,'',3,3,'Dance/Electronic','files/Flume/Skin/01 Helix.mp3',0),
(4,'It''s Just Us Now',1,'',4,4,'Alternative/Indie','files/Sleigh Bells/Jessica Rabbit/01 It''s Just Us Now.mp3',0),
(5,'Wilderness',1,'',5,5,'Alternative/Indie','files/Explosions In The Sky/The Wilderness/01 Wilderness.mp3',2016),
(6,'Fuck With Myself',2,'',2,2,'Alternative/Indie','files/BANKS/The Altar/02 Fuck With Myself.mp3',2016),
(7,'Never Be Like You (feat. Kai)',2,'',3,3,'Dance/Electronic','files/Flume/Skin/02 Never Be Like You (feat_ Kai).mp3',0),
(8,'The Ecstatics',2,'',5,5,'Alternative/Indie','files/Explosions In The Sky/The Wilderness/02 The Ecstatics.mp3',0),
(9,'Torn Clean',2,'',4,4,'Alternative/Indie','files/Sleigh Bells/Jessica Rabbit/02 Torn Clean.mp3',2016),
(10,'We Had Everything',2,'',1,1,'Alternative/Indie','files/Broods/Conscious/02 We Had Everything.mp3',2016),
(11,'Funeral Pyre',1,'',6,6,'Alternative/Indie','files/Phantogram/Three/01 Funeral Pyre.mp3',2016),
(12,'Same Old Blues',2,'',6,6,'Alternative/Indie','files/Phantogram/Three/02 Same Old Blues.mp3',2016);


-- select s.id, s.title, s.track, s.comment, s.year, s.last_played, s.path, s.album_id, s.artist_id, al.title as album_title, ar.name from songs as s
-- full join albums as al on s.album_id = al.id
-- full join artists as ar on s.artist_id = ar.id;