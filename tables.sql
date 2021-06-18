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

create table genres (
  name varchar primary key UNIQUE
  );
