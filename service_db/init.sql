--
-- PostgreSQL database dump
--

-- Dumped from database version 15.10 (Debian 15.10-0+deb12u1)
-- Dumped by pg_dump version 17.0

-- Started on 2024-11-23 01:09:43

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE IF EXISTS dunice;
--
-- TOC entry 3398 (class 1262 OID 16388)
-- Name: dunice; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE dunice WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.UTF-8';


ALTER DATABASE dunice OWNER TO postgres;

\connect dunice

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 6 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 3400 (class 0 OID 0)
-- Dependencies: 6
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 222 (class 1259 OID 16431)
-- Name: admins; Type: TABLE; Schema: public; Owner: dunice_user
--

CREATE TABLE public.admins (
    userid integer NOT NULL,
    adminid integer NOT NULL
);


ALTER TABLE public.admins OWNER TO dunice_user;

--
-- TOC entry 221 (class 1259 OID 16430)
-- Name: admins_adminid_seq; Type: SEQUENCE; Schema: public; Owner: dunice_user
--

CREATE SEQUENCE public.admins_adminid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.admins_adminid_seq OWNER TO dunice_user;

--
-- TOC entry 3402 (class 0 OID 0)
-- Dependencies: 221
-- Name: admins_adminid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dunice_user
--

ALTER SEQUENCE public.admins_adminid_seq OWNED BY public.admins.adminid;


--
-- TOC entry 219 (class 1259 OID 16409)
-- Name: offices; Type: TABLE; Schema: public; Owner: dunice_user
--

CREATE TABLE public.offices (
    officeid integer NOT NULL,
    address character varying(255) NOT NULL
);


ALTER TABLE public.offices OWNER TO dunice_user;

--
-- TOC entry 218 (class 1259 OID 16408)
-- Name: offices_officeid_seq; Type: SEQUENCE; Schema: public; Owner: dunice_user
--

CREATE SEQUENCE public.offices_officeid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.offices_officeid_seq OWNER TO dunice_user;

--
-- TOC entry 3403 (class 0 OID 0)
-- Dependencies: 218
-- Name: offices_officeid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dunice_user
--

ALTER SEQUENCE public.offices_officeid_seq OWNED BY public.offices.officeid;


--
-- TOC entry 220 (class 1259 OID 16415)
-- Name: user_offices; Type: TABLE; Schema: public; Owner: dunice_user
--

CREATE TABLE public.user_offices (
    userid integer NOT NULL,
    officeid integer NOT NULL
);


ALTER TABLE public.user_offices OWNER TO dunice_user;

--
-- TOC entry 217 (class 1259 OID 16398)
-- Name: users; Type: TABLE; Schema: public; Owner: dunice_user
--

CREATE TABLE public.users (
    id integer NOT NULL,
    login character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    name character varying(100),
    surname character varying(100),
    middlename character varying(100)
);


ALTER TABLE public.users OWNER TO dunice_user;

--
-- TOC entry 216 (class 1259 OID 16397)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: dunice_user
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO dunice_user;

--
-- TOC entry 3404 (class 0 OID 0)
-- Dependencies: 216
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dunice_user
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 224 (class 1259 OID 16443)
-- Name: workspaces; Type: TABLE; Schema: public; Owner: dunice_user
--

CREATE TABLE public.workspaces (
    userid integer NOT NULL,
    workspaceid integer NOT NULL
);


ALTER TABLE public.workspaces OWNER TO dunice_user;

--
-- TOC entry 223 (class 1259 OID 16442)
-- Name: workspaces_workspaceid_seq; Type: SEQUENCE; Schema: public; Owner: dunice_user
--

CREATE SEQUENCE public.workspaces_workspaceid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.workspaces_workspaceid_seq OWNER TO dunice_user;

--
-- TOC entry 3405 (class 0 OID 0)
-- Dependencies: 223
-- Name: workspaces_workspaceid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dunice_user
--

ALTER SEQUENCE public.workspaces_workspaceid_seq OWNED BY public.workspaces.workspaceid;


--
-- TOC entry 3224 (class 2604 OID 16434)
-- Name: admins adminid; Type: DEFAULT; Schema: public; Owner: dunice_user
--

ALTER TABLE ONLY public.admins ALTER COLUMN adminid SET DEFAULT nextval('public.admins_adminid_seq'::regclass);


--
-- TOC entry 3223 (class 2604 OID 16412)
-- Name: offices officeid; Type: DEFAULT; Schema: public; Owner: dunice_user
--

ALTER TABLE ONLY public.offices ALTER COLUMN officeid SET DEFAULT nextval('public.offices_officeid_seq'::regclass);


--
-- TOC entry 3222 (class 2604 OID 16401)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: dunice_user
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3225 (class 2604 OID 16446)
-- Name: workspaces workspaceid; Type: DEFAULT; Schema: public; Owner: dunice_user
--

ALTER TABLE ONLY public.workspaces ALTER COLUMN workspaceid SET DEFAULT nextval('public.workspaces_workspaceid_seq'::regclass);


--
-- TOC entry 3390 (class 0 OID 16431)
-- Dependencies: 222
-- Data for Name: admins; Type: TABLE DATA; Schema: public; Owner: dunice_user
--



--
-- TOC entry 3387 (class 0 OID 16409)
-- Dependencies: 219
-- Data for Name: offices; Type: TABLE DATA; Schema: public; Owner: dunice_user
--



--
-- TOC entry 3388 (class 0 OID 16415)
-- Dependencies: 220
-- Data for Name: user_offices; Type: TABLE DATA; Schema: public; Owner: dunice_user
--



--
-- TOC entry 3385 (class 0 OID 16398)
-- Dependencies: 217
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: dunice_user
--



--
-- TOC entry 3392 (class 0 OID 16443)
-- Dependencies: 224
-- Data for Name: workspaces; Type: TABLE DATA; Schema: public; Owner: dunice_user
--



--
-- TOC entry 3406 (class 0 OID 0)
-- Dependencies: 221
-- Name: admins_adminid_seq; Type: SEQUENCE SET; Schema: public; Owner: dunice_user
--

SELECT pg_catalog.setval('public.admins_adminid_seq', 1, false);


--
-- TOC entry 3407 (class 0 OID 0)
-- Dependencies: 218
-- Name: offices_officeid_seq; Type: SEQUENCE SET; Schema: public; Owner: dunice_user
--

SELECT pg_catalog.setval('public.offices_officeid_seq', 1, false);


--
-- TOC entry 3408 (class 0 OID 0)
-- Dependencies: 216
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dunice_user
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- TOC entry 3409 (class 0 OID 0)
-- Dependencies: 223
-- Name: workspaces_workspaceid_seq; Type: SEQUENCE SET; Schema: public; Owner: dunice_user
--

SELECT pg_catalog.setval('public.workspaces_workspaceid_seq', 1, false);


--
-- TOC entry 3235 (class 2606 OID 16436)
-- Name: admins admins_pkey; Type: CONSTRAINT; Schema: public; Owner: dunice_user
--

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_pkey PRIMARY KEY (adminid);


--
-- TOC entry 3231 (class 2606 OID 16414)
-- Name: offices offices_pkey; Type: CONSTRAINT; Schema: public; Owner: dunice_user
--

ALTER TABLE ONLY public.offices
    ADD CONSTRAINT offices_pkey PRIMARY KEY (officeid);


--
-- TOC entry 3233 (class 2606 OID 16419)
-- Name: user_offices user_offices_pkey; Type: CONSTRAINT; Schema: public; Owner: dunice_user
--

ALTER TABLE ONLY public.user_offices
    ADD CONSTRAINT user_offices_pkey PRIMARY KEY (userid, officeid);


--
-- TOC entry 3227 (class 2606 OID 16407)
-- Name: users users_login_key; Type: CONSTRAINT; Schema: public; Owner: dunice_user
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_login_key UNIQUE (login);


--
-- TOC entry 3229 (class 2606 OID 16405)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: dunice_user
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3237 (class 2606 OID 16448)
-- Name: workspaces workspaces_pkey; Type: CONSTRAINT; Schema: public; Owner: dunice_user
--

ALTER TABLE ONLY public.workspaces
    ADD CONSTRAINT workspaces_pkey PRIMARY KEY (workspaceid);


--
-- TOC entry 3240 (class 2606 OID 16437)
-- Name: admins admins_userid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dunice_user
--

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_userid_fkey FOREIGN KEY (userid) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- TOC entry 3238 (class 2606 OID 16425)
-- Name: user_offices user_offices_officeid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dunice_user
--

ALTER TABLE ONLY public.user_offices
    ADD CONSTRAINT user_offices_officeid_fkey FOREIGN KEY (officeid) REFERENCES public.offices(officeid) ON DELETE CASCADE;


--
-- TOC entry 3239 (class 2606 OID 16420)
-- Name: user_offices user_offices_userid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dunice_user
--

ALTER TABLE ONLY public.user_offices
    ADD CONSTRAINT user_offices_userid_fkey FOREIGN KEY (userid) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- TOC entry 3241 (class 2606 OID 16449)
-- Name: workspaces workspaces_userid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dunice_user
--

ALTER TABLE ONLY public.workspaces
    ADD CONSTRAINT workspaces_userid_fkey FOREIGN KEY (userid) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- TOC entry 3399 (class 0 OID 0)
-- Dependencies: 3398
-- Name: DATABASE dunice; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON DATABASE dunice TO dunice_user;


--
-- TOC entry 3401 (class 0 OID 0)
-- Dependencies: 6
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: pg_database_owner
--

GRANT ALL ON SCHEMA public TO dunice_user;


-- Completed on 2024-11-23 01:09:55

--
-- PostgreSQL database dump complete
--

