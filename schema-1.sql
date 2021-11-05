--
-- PostgreSQL database dump
--

-- Dumped from database version 12.7 (Ubuntu 12.7-1.pgdg20.04+1)
-- Dumped by pg_dump version 12.7 (Ubuntu 12.7-1.pgdg20.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: albums; Type: TABLE; Schema: public; Owner: zach
--

CREATE TABLE public.albums (
    id integer NOT NULL,
    title character varying,
    artist_id integer,
    image boolean DEFAULT false
);


ALTER TABLE public.albums OWNER TO zach;

--
-- Name: albums_id_seq; Type: SEQUENCE; Schema: public; Owner: zach
--

CREATE SEQUENCE public.albums_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.albums_id_seq OWNER TO zach;

--
-- Name: albums_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zach
--

ALTER SEQUENCE public.albums_id_seq OWNED BY public.albums.id;


--
-- Name: artists; Type: TABLE; Schema: public; Owner: zach
--

CREATE TABLE public.artists (
    id integer NOT NULL,
    name character varying
);


ALTER TABLE public.artists OWNER TO zach;

--
-- Name: artists_id_seq; Type: SEQUENCE; Schema: public; Owner: zach
--

CREATE SEQUENCE public.artists_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.artists_id_seq OWNER TO zach;

--
-- Name: artists_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zach
--

ALTER SEQUENCE public.artists_id_seq OWNED BY public.artists.id;


--
-- Name: genres; Type: TABLE; Schema: public; Owner: zach
--

CREATE TABLE public.genres (
    name character varying NOT NULL
);


ALTER TABLE public.genres OWNER TO zach;

--
-- Name: liked_songs; Type: TABLE; Schema: public; Owner: zach
--

CREATE TABLE public.liked_songs (
    id integer NOT NULL,
    song_id integer NOT NULL,
    date_added timestamp without time zone NOT NULL
);


ALTER TABLE public.liked_songs OWNER TO zach;

--
-- Name: liked_songs_id_seq; Type: SEQUENCE; Schema: public; Owner: zach
--

CREATE SEQUENCE public.liked_songs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.liked_songs_id_seq OWNER TO zach;

--
-- Name: liked_songs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zach
--

ALTER SEQUENCE public.liked_songs_id_seq OWNED BY public.liked_songs.id;


--
-- Name: songs; Type: TABLE; Schema: public; Owner: zach
--

CREATE TABLE public.songs (
    id integer NOT NULL,
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
    duration numeric
);


ALTER TABLE public.songs OWNER TO zach;

--
-- Name: songs_id_seq; Type: SEQUENCE; Schema: public; Owner: zach
--

CREATE SEQUENCE public.songs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.songs_id_seq OWNER TO zach;

--
-- Name: songs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zach
--

ALTER SEQUENCE public.songs_id_seq OWNED BY public.songs.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: zach
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying NOT NULL,
    password character varying NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO zach;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: zach
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO zach;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zach
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: albums id; Type: DEFAULT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.albums ALTER COLUMN id SET DEFAULT nextval('public.albums_id_seq'::regclass);


--
-- Name: artists id; Type: DEFAULT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.artists ALTER COLUMN id SET DEFAULT nextval('public.artists_id_seq'::regclass);


--
-- Name: liked_songs id; Type: DEFAULT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.liked_songs ALTER COLUMN id SET DEFAULT nextval('public.liked_songs_id_seq'::regclass);


--
-- Name: songs id; Type: DEFAULT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.songs ALTER COLUMN id SET DEFAULT nextval('public.songs_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: albums albums_pkey; Type: CONSTRAINT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.albums
    ADD CONSTRAINT albums_pkey PRIMARY KEY (id);


--
-- Name: albums albums_title_artist_id_key; Type: CONSTRAINT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.albums
    ADD CONSTRAINT albums_title_artist_id_key UNIQUE (title, artist_id);


--
-- Name: artists artists_name_key; Type: CONSTRAINT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.artists
    ADD CONSTRAINT artists_name_key UNIQUE (name);


--
-- Name: artists artists_pkey; Type: CONSTRAINT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.artists
    ADD CONSTRAINT artists_pkey PRIMARY KEY (id);


--
-- Name: genres genres_pkey; Type: CONSTRAINT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.genres
    ADD CONSTRAINT genres_pkey PRIMARY KEY (name);


--
-- Name: songs songs_pkey; Type: CONSTRAINT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.songs
    ADD CONSTRAINT songs_pkey PRIMARY KEY (id);


--
-- Name: songs songs_title_artist_id_album_id_key; Type: CONSTRAINT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.songs
    ADD CONSTRAINT songs_title_artist_id_album_id_key UNIQUE (title, artist_id, album_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: albums albums_artist_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.albums
    ADD CONSTRAINT albums_artist_id_fkey FOREIGN KEY (artist_id) REFERENCES public.artists(id);


--
-- Name: liked_songs liked_songs_song_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.liked_songs
    ADD CONSTRAINT liked_songs_song_id_fkey FOREIGN KEY (song_id) REFERENCES public.songs(id);


--
-- Name: songs songs_album_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.songs
    ADD CONSTRAINT songs_album_id_fkey FOREIGN KEY (album_id) REFERENCES public.albums(id);


--
-- Name: songs songs_artist_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.songs
    ADD CONSTRAINT songs_artist_id_fkey FOREIGN KEY (artist_id) REFERENCES public.artists(id);


--
-- Name: songs songs_genre_fkey; Type: FK CONSTRAINT; Schema: public; Owner: zach
--

ALTER TABLE ONLY public.songs
    ADD CONSTRAINT songs_genre_fkey FOREIGN KEY (genre) REFERENCES public.genres(name);


--
-- PostgreSQL database dump complete
--

