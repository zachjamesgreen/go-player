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

INSERT INTO artists (id,name) VALUES
(1, 'Broods'),
(2, 'BANKS'),
(3, 'Flume'),
(4, 'Sleigh Bells'),
(5, 'Explosions In The Sky');

INSERT INTO albums (id,title,artist_id) VALUES
(1, 'Conscious', 1),
(2, 'The Altar', 2),
(3, 'Skin', 3),
(4, 'Jessica Rabbit', 4),
(5, 'The Wilderness', 5);

INSERT INTO genres (name) VALUES
('Dance/Electronic'),
('Alternative/Indie');

INSERT INTO songs (id,title,track,comment,album_id,artist_id,genre,path) VALUES
(1,'Free',1,'',1,1,'Alternative/Indie','files/Broods/Conscious'),
(2,'Gemini Feed',1,'',2,2,'Alternative/Indie','files/BANKS/The Altar'),
(3,'Helix',1,'',3,3,'Dance/Electronic','files/Flume/Skin'),
(4,'It''s Just Us Now',1,'',4,4,'Alternative/Indie','files/Sleigh Bells/Jessica Rabbit'),
(5,'Wilderness',1,'',5,5,'Alternative/Indie','files/Explosions In The Sky/The Wilderness'),
(6,'Fuck With Myself',2,'',2,2,'Alternative/Indie','files/BANKS/The Altar'),
(7,'Never Be Like You (feat. Kai)',2,'',3,3,'Dance/Electronic','files/Flume/Skin'),
(8,'The Ecstatics',2,'',5,5,'Alternative/Indie','files/Explosions In The Sky/The Wilderness'),
(9,'Torn Clean',2,'',4,4,'Alternative/Indie','files/Sleigh Bells/Jessica Rabbit'),
(10,'We Had Everything',2,'',1,1,'Alternative/Indie','files/Broods/Conscious');