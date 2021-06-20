create table artists (
  id serial primary key,
  name varchar UNIQUE
  );

create table albums (
  id serial primary key,
  title varchar,
  artist_id integer references artists(id),
  UNIQUE (title, artist_id)
  );

create table genres (
  name varchar primary key UNIQUE
  );

create table songs (
  id serial primary key,
  title varchar,
  track int,
  comment text,
  album_id integer references albums(id),
  artist_id integer references artists(id),
  genre varchar references genres(name),
  path varchar,
  UNIQUE (title, artist_id, album_id)
  );

ALTER SEQUENCE albums_id_seq RESTART WITH 1;
ALTER SEQUENCE artists_id_seq RESTART WITH 1;
ALTER SEQUENCE songs_id_seq RESTART WITH 1;


COPY public.artists (id, name) FROM stdin;
2	BANKS
3	Flume
5	Explosions In The Sky
4	Sleigh Bells
1	Broods
\.

COPY public.albums (id, title, artist_id) FROM stdin;
2	The Altar	2
3	Skin	3
5	The Wilderness	5
4	Jessica Rabbit	4
1	Conscious	1
\.

COPY public.genres (name) FROM stdin;
Dance/Electronic
Alternative/Indie
\.

COPY public.songs (id, title, track, comment, album_id, artist_id, genre, path) FROM stdin;
1	Free	1		1	1	Alternative/Indie	files/Broods/Conscious
2	Gemini Feed	1		2	2	Alternative/Indie	files/BANKS/The Altar
3	Helix	1		3	3	Dance/Electronic	files/Flume/Skin
4	It's Just Us Now	1		4	4	Alternative/Indie	files/Sleigh Bells/Jessica Rabbit
5	Wilderness	1		5	5	Alternative/Indie	files/Explosions In The Sky/The Wilderness
6	Fuck With Myself	2		2	2	Alternative/Indie	files/BANKS/The Altar
7	Never Be Like You (feat. Kai)	2		3	3	Dance/Electronic	files/Flume/Skin
8	The Ecstatics	2		5	5	Alternative/Indie	files/Explosions In The Sky/The Wilderness
9	Torn Clean	2		4	4	Alternative/Indie	files/Sleigh Bells/Jessica Rabbit
10	We Had Everything	2		1	1	Alternative/Indie	files/Broods/Conscious
\.