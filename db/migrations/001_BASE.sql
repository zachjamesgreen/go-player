-- +goose Up

CREATE TABLE public.artists (
    id SERIAL PRIMARY KEY,
    name character varying,
    spotify_id character varying DEFAULT ''::character varying,
    images json DEFAULT '{}'::json
);

CREATE TABLE public.albums (
    id SERIAL PRIMARY KEY,
    title character varying,
    artist_id integer,
    image boolean DEFAULT false,
    spotify_id character varying DEFAULT ''::character varying,
    spotify_link character varying DEFAULT ''::character varying,
    images json DEFAULT '{}'::json
);

CREATE TABLE public.songs (
    id SERIAL PRIMARY KEY,
    title character varying,
    track integer,
    comment text,
    album_id integer,
    artist_id integer,
    genre character varying,
    path character varying DEFAULT ''::character varying,
    last_played timestamp without time zone,
    year integer DEFAULT 0,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    duration numeric,
    liked boolean DEFAULT false,
    liked_date timestamp without time zone DEFAULT now()
);

CREATE TABLE public.users (
    id SERIAL PRIMARY KEY,
    username character varying NOT NULL,
    password character varying NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

ALTER TABLE ONLY public.albums ADD CONSTRAINT albums_title_artist_id_key UNIQUE (title, artist_id);
ALTER TABLE ONLY public.artists ADD CONSTRAINT artists_name_key UNIQUE (name);
ALTER TABLE ONLY public.songs ADD CONSTRAINT songs_title_artist_id_album_id_key UNIQUE (title, artist_id, album_id);
ALTER TABLE ONLY public.users ADD CONSTRAINT users_username_key UNIQUE (username);
ALTER TABLE ONLY public.albums ADD CONSTRAINT albums_artist_id_fkey FOREIGN KEY (artist_id) REFERENCES public.artists(id);
ALTER TABLE ONLY public.songs ADD CONSTRAINT songs_album_id_fkey FOREIGN KEY (album_id) REFERENCES public.albums(id);
ALTER TABLE ONLY public.songs ADD CONSTRAINT songs_artist_id_fkey FOREIGN KEY (artist_id) REFERENCES public.artists(id);
-- +goose Down
DROP TABLE public.songs;
DROP TABLE public.albums;
DROP TABLE public.artists;
DROP TABLE public.users;