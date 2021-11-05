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
-- Data for Name: albums; Type: TABLE DATA; Schema: public; Owner: zach
--

COPY public.albums (id, title, artist_id, image) FROM stdin;
1	Conscious	1	t
2	The Altar	2	t
3	Skin	3	t
4	Jessica Rabbit	4	t
5	The Wilderness	5	t
6	Phantogram	6	t
9	Three	9	t
10	1000 Forms of Fear	10	f
21	Love Death Immortality	21	f
23	House of Balloons	23	f
24	WOMB	24	f
\.


--
-- Data for Name: artists; Type: TABLE DATA; Schema: public; Owner: zach
--

COPY public.artists (id, name) FROM stdin;
6	Three
1	Broods
2	BANKS
3	Flume
4	Sleigh Bells
5	Explosions In The Sky
9	Phantogram
10	Sia
21	The Glitch Mob
23	The Weeknd
24	Purity Ring
\.


--
-- Data for Name: genres; Type: TABLE DATA; Schema: public; Owner: zach
--

COPY public.genres (name) FROM stdin;
Dance/Electronic
Alternative/Indie
Pop
[Pop
Electronic

\.


--
-- Data for Name: liked_songs; Type: TABLE DATA; Schema: public; Owner: zach
--

COPY public.liked_songs (id, song_id, date_added) FROM stdin;
1	8	2021-07-12 17:22:04.465207
2	9	2021-07-12 18:13:12.673939
3	9	2021-07-12 18:18:06.133031
\.


--
-- Data for Name: songs; Type: TABLE DATA; Schema: public; Owner: zach
--

COPY public.songs (id, title, track, comment, album_id, artist_id, genre, path, last_played, year, created_at, updated_at, duration) FROM stdin;
8	The Ecstatics	2		5	5	Alternative/Indie	files/Explosions In The Sky/The Wilderness/02 The Ecstatics.mp3	\N	0	2021-06-25 13:51:00.80028	2021-06-25 13:51:16.85101	0
9	Torn Clean	2		4	4	Alternative/Indie	files/Sleigh Bells/Jessica Rabbit/02 Torn Clean.mp3	\N	2016	2021-06-25 13:51:00.80028	2021-06-25 13:51:16.85101	0
16	Big Girls Cry	2		10	10	[Pop	files/Sia/1000 Forms of Fear/02 Big Girls Cry.mp3	\N	0	2021-06-25 16:08:43.704859	2021-06-25 16:08:43.704859	0
26	Beauty of the Unhidden Heart (feat. Sister Crayon)	10		21	21	Electronic	files/The Glitch Mob/Love Death Immortality/Beauty of the Unhidden Heart (feat. Sister Crayon).mp3	\N	0	2021-07-11 00:56:41.689885	2021-07-11 00:56:41.689885	0
27	House Of Balloons / Glass Table Girls	3		23	23		files/The Weeknd/House of Balloons/House Of Balloons _ Glass Table Girls.mp3	\N	0	2021-07-11 00:58:03.023393	2021-07-11 00:58:03.023393	0
25	Can't Kill Us	5		21	21	Electronic	files/The Glitch Mob/Love Death Immortality/Can't Kill Us.mp3	\N	0	2021-07-11 00:54:47.009147	2021-07-11 00:54:47.009147	0
10	We Had Everything	2		1	1	Alternative/Indie	files/Broods/Conscious/02 We Had Everything.mp3	\N	2016	2021-06-25 13:51:00.80028	2021-06-25 13:51:16.85101	0
11	Funeral Pyre	1		6	6	Alternative/Indie	files/Phantogram/Three/01 Funeral Pyre.mp3	\N	2016	2021-06-25 13:51:00.80028	2021-06-25 13:51:16.85101	0
12	Same Old Blues	2		6	6	Alternative/Indie	files/Phantogram/Three/02 Same Old Blues.mp3	\N	2016	2021-06-25 13:51:00.80028	2021-06-25 13:51:16.85101	0
29	almanac	9		24	24		files/Purity Ring/WOMB/almanac.flac	\N	2020	2021-07-11 01:14:54.920476	2021-07-11 01:14:54.920476	0
1	Free	1		1	1	Alternative/Indie	files/Broods/Conscious/01 Free.mp3	\N	2016	2021-06-25 13:51:00.80028	2021-06-25 13:51:16.85101	0
2	Gemini Feed	1		2	2	Alternative/Indie	files/BANKS/The Altar/01 Gemini Feed.mp3	\N	2016	2021-06-25 13:51:00.80028	2021-06-25 13:51:16.85101	0
3	Helix	1		3	3	Dance/Electronic	files/Flume/Skin/01 Helix.mp3	\N	0	2021-06-25 13:51:00.80028	2021-06-25 13:51:16.85101	0
4	It's Just Us Now	1		4	4	Alternative/Indie	files/Sleigh Bells/Jessica Rabbit/01 It's Just Us Now.mp3	\N	0	2021-06-25 13:51:00.80028	2021-06-25 13:51:16.85101	0
5	Wilderness	1		5	5	Alternative/Indie	files/Explosions In The Sky/The Wilderness/01 Wilderness.mp3	\N	2016	2021-06-25 13:51:00.80028	2021-06-25 13:51:16.85101	0
15	Chandelier	1		10	10	Pop	files/Sia/1000 Forms of Fear/01 Chandelier.mp3	\N	0	2021-06-25 16:08:43.654579	2021-06-25 16:08:43.654579	0
28	stardew	10		24	24		files/Purity Ring/WOMB/stardew.flac	\N	2020	2021-07-11 01:10:43.789167	2021-07-11 01:10:43.789167	0
6	Fuck With Myself	2		2	2	Alternative/Indie	files/BANKS/The Altar/02 Fuck With Myself.mp3	\N	2016	2021-06-25 13:51:00.80028	2021-06-25 13:51:16.85101	0
7	Never Be Like You (feat. Kai)	2		3	3	Dance/Electronic	files/Flume/Skin/02 Never Be Like You (feat_ Kai).mp3	\N	0	2021-06-25 13:51:00.80028	2021-06-25 13:51:16.85101	0
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: zach
--

COPY public.users (id, username, password, created_at, updated_at) FROM stdin;
1	zach	$2a$14$hEDv.Nx6mkCLVPmGB7e6QOc3gDvrthATsUf8heALp.5FHTPOmrTpi	2021-06-24 12:52:47.094424	2021-06-24 12:52:47.094424
\.


--
-- Name: albums_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zach
--

SELECT pg_catalog.setval('public.albums_id_seq', 26, true);


--
-- Name: artists_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zach
--

SELECT pg_catalog.setval('public.artists_id_seq', 26, true);


--
-- Name: liked_songs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zach
--

SELECT pg_catalog.setval('public.liked_songs_id_seq', 3, true);


--
-- Name: songs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zach
--

SELECT pg_catalog.setval('public.songs_id_seq', 29, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zach
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


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

