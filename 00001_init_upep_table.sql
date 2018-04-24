-- +goose Up
-- SQL in this section is executed when the migration is applied.
--
-- PostgreSQL database dump
--

-- Dumped from database version 10.3
-- Dumped by pg_dump version 10.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: upep; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA upep;


ALTER SCHEMA upep OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: upep_accessions; Type: TABLE; Schema: upep; Owner: postgres
--

CREATE TABLE upep.upep_accessions (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    accession text NOT NULL
);


ALTER TABLE upep.upep_accessions OWNER TO postgres;

--
-- Name: upep_accessions_id_seq; Type: SEQUENCE; Schema: upep; Owner: postgres
--

CREATE SEQUENCE upep.upep_accessions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE upep.upep_accessions_id_seq OWNER TO postgres;

--
-- Name: upep_accessions_id_seq; Type: SEQUENCE OWNED BY; Schema: upep; Owner: postgres
--

ALTER SEQUENCE upep.upep_accessions_id_seq OWNED BY upep.upep_accessions.id;


--
-- Name: upep_blast_db; Type: TABLE; Schema: upep; Owner: postgres
--

CREATE TABLE upep.upep_blast_db (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    title text NOT NULL,
    path text NOT NULL,
    description text NOT NULL
);


ALTER TABLE upep.upep_blast_db OWNER TO postgres;

--
-- Name: upep_features; Type: TABLE; Schema: upep; Owner: postgres
--

CREATE TABLE upep.upep_features (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    feature_start integer NOT NULL,
    feature_end integer NOT NULL,
    name character varying(10) NOT NULL,
    partial_start boolean DEFAULT false NOT NULL,
    partial_end boolean DEFAULT false NOT NULL,
    ref_seq_entry_id bigint NOT NULL
);


ALTER TABLE upep.upep_features OWNER TO postgres;

--
-- Name: upep_features_id_seq; Type: SEQUENCE; Schema: upep; Owner: postgres
--

CREATE SEQUENCE upep.upep_features_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE upep.upep_features_id_seq OWNER TO postgres;

--
-- Name: upep_features_id_seq; Type: SEQUENCE OWNED BY; Schema: upep; Owner: postgres
--

ALTER SEQUENCE upep.upep_features_id_seq OWNED BY upep.upep_features.id;


--
-- Name: upep_gene_identifiers; Type: TABLE; Schema: upep; Owner: postgres
--

CREATE TABLE upep.upep_gene_identifiers (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    gi bigint NOT NULL
);


ALTER TABLE upep.upep_gene_identifiers OWNER TO postgres;

--
-- Name: upep_gis_id_seq; Type: SEQUENCE; Schema: upep; Owner: postgres
--

CREATE SEQUENCE upep.upep_gis_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE upep.upep_gis_id_seq OWNER TO postgres;

--
-- Name: upep_gis_id_seq; Type: SEQUENCE OWNED BY; Schema: upep; Owner: postgres
--

ALTER SEQUENCE upep.upep_gis_id_seq OWNED BY upep.upep_gene_identifiers.id;


--
-- Name: upep_molecular_types; Type: TABLE; Schema: upep; Owner: postgres
--

CREATE TABLE upep.upep_molecular_types (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name character varying(10) NOT NULL
);


ALTER TABLE upep.upep_molecular_types OWNER TO postgres;

--
-- Name: upep_molecular_types_id_seq; Type: SEQUENCE; Schema: upep; Owner: postgres
--

CREATE SEQUENCE upep.upep_molecular_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE upep.upep_molecular_types_id_seq OWNER TO postgres;

--
-- Name: upep_molecular_types_id_seq; Type: SEQUENCE OWNED BY; Schema: upep; Owner: postgres
--

ALTER SEQUENCE upep.upep_molecular_types_id_seq OWNED BY upep.upep_molecular_types.id;


--
-- Name: upep_organisms; Type: TABLE; Schema: upep; Owner: postgres
--

CREATE TABLE upep.upep_organisms (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name text NOT NULL
);


ALTER TABLE upep.upep_organisms OWNER TO postgres;

--
-- Name: upep_organisms_id_seq; Type: SEQUENCE; Schema: upep; Owner: postgres
--

CREATE SEQUENCE upep.upep_organisms_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE upep.upep_organisms_id_seq OWNER TO postgres;

--
-- Name: upep_organisms_id_seq; Type: SEQUENCE OWNED BY; Schema: upep; Owner: postgres
--

ALTER SEQUENCE upep.upep_organisms_id_seq OWNED BY upep.upep_organisms.id;


--
-- Name: upep_ref_seq_db; Type: TABLE; Schema: upep; Owner: postgres
--

CREATE TABLE upep.upep_ref_seq_db (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name text NOT NULL,
    version integer DEFAULT 0 NOT NULL
);


ALTER TABLE upep.upep_ref_seq_db OWNER TO postgres;

--
-- Name: upep_ref_seq_db_id_seq; Type: SEQUENCE; Schema: upep; Owner: postgres
--

CREATE SEQUENCE upep.upep_ref_seq_db_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE upep.upep_ref_seq_db_id_seq OWNER TO postgres;

--
-- Name: upep_ref_seq_db_id_seq; Type: SEQUENCE OWNED BY; Schema: upep; Owner: postgres
--

ALTER SEQUENCE upep.upep_ref_seq_db_id_seq OWNED BY upep.upep_ref_seq_db.id;


--
-- Name: upep_ref_seq_entries; Type: TABLE; Schema: upep; Owner: postgres
--

CREATE TABLE upep.upep_ref_seq_entries (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name character varying(20),
    organism_id bigint,
    molecular_type_id bigint,
    accession_id bigint,
    gi_id bigint,
    ref_seq_db_id bigint,
    ref_seq_sequence text NOT NULL
);


ALTER TABLE upep.upep_ref_seq_entries OWNER TO postgres;

--
-- Name: upep_ref_seq_entries_id_seq; Type: SEQUENCE; Schema: upep; Owner: postgres
--

CREATE SEQUENCE upep.upep_ref_seq_entries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE upep.upep_ref_seq_entries_id_seq OWNER TO postgres;

--
-- Name: upep_ref_seq_entries_id_seq; Type: SEQUENCE OWNED BY; Schema: upep; Owner: postgres
--

ALTER SEQUENCE upep.upep_ref_seq_entries_id_seq OWNED BY upep.upep_ref_seq_entries.id;


--
-- Name: upep_accessions id; Type: DEFAULT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_accessions ALTER COLUMN id SET DEFAULT nextval('upep.upep_accessions_id_seq'::regclass);


--
-- Name: upep_features id; Type: DEFAULT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_features ALTER COLUMN id SET DEFAULT nextval('upep.upep_features_id_seq'::regclass);


--
-- Name: upep_gene_identifiers id; Type: DEFAULT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_gene_identifiers ALTER COLUMN id SET DEFAULT nextval('upep.upep_gis_id_seq'::regclass);


--
-- Name: upep_molecular_types id; Type: DEFAULT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_molecular_types ALTER COLUMN id SET DEFAULT nextval('upep.upep_molecular_types_id_seq'::regclass);


--
-- Name: upep_organisms id; Type: DEFAULT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_organisms ALTER COLUMN id SET DEFAULT nextval('upep.upep_organisms_id_seq'::regclass);


--
-- Name: upep_ref_seq_db id; Type: DEFAULT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_ref_seq_db ALTER COLUMN id SET DEFAULT nextval('upep.upep_ref_seq_db_id_seq'::regclass);


--
-- Name: upep_ref_seq_entries id; Type: DEFAULT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_ref_seq_entries ALTER COLUMN id SET DEFAULT nextval('upep.upep_ref_seq_entries_id_seq'::regclass);


--
-- Data for Name: upep_accessions; Type: TABLE DATA; Schema: upep; Owner: postgres
--

COPY upep.upep_accessions (id, created_at, updated_at, accession) FROM stdin;
\.


--
-- Data for Name: upep_blast_db; Type: TABLE DATA; Schema: upep; Owner: postgres
--

COPY upep.upep_blast_db (id, created_at, updated_at, title, path, description) FROM stdin;
\.


--
-- Data for Name: upep_features; Type: TABLE DATA; Schema: upep; Owner: postgres
--

COPY upep.upep_features (id, created_at, updated_at, feature_start, feature_end, name, partial_start, partial_end, ref_seq_entry_id) FROM stdin;
\.


--
-- Data for Name: upep_gene_identifiers; Type: TABLE DATA; Schema: upep; Owner: postgres
--

COPY upep.upep_gene_identifiers (id, created_at, updated_at, gi) FROM stdin;
\.


--
-- Data for Name: upep_molecular_types; Type: TABLE DATA; Schema: upep; Owner: postgres
--

COPY upep.upep_molecular_types (id, created_at, updated_at, name) FROM stdin;
\.


--
-- Data for Name: upep_organisms; Type: TABLE DATA; Schema: upep; Owner: postgres
--

COPY upep.upep_organisms (id, created_at, updated_at, name) FROM stdin;
\.


--
-- Data for Name: upep_ref_seq_db; Type: TABLE DATA; Schema: upep; Owner: postgres
--

COPY upep.upep_ref_seq_db (id, created_at, updated_at, name, version) FROM stdin;
\.


--
-- Data for Name: upep_ref_seq_entries; Type: TABLE DATA; Schema: upep; Owner: postgres
--

COPY upep.upep_ref_seq_entries (id, created_at, updated_at, name, organism_id, molecular_type_id, accession_id, gi_id, ref_seq_db_id, ref_seq_sequence) FROM stdin;
\.


--
-- Name: upep_accessions_id_seq; Type: SEQUENCE SET; Schema: upep; Owner: postgres
--

SELECT pg_catalog.setval('upep.upep_accessions_id_seq', 1, false);


--
-- Name: upep_features_id_seq; Type: SEQUENCE SET; Schema: upep; Owner: postgres
--

SELECT pg_catalog.setval('upep.upep_features_id_seq', 1, false);


--
-- Name: upep_gis_id_seq; Type: SEQUENCE SET; Schema: upep; Owner: postgres
--

SELECT pg_catalog.setval('upep.upep_gis_id_seq', 1, false);


--
-- Name: upep_molecular_types_id_seq; Type: SEQUENCE SET; Schema: upep; Owner: postgres
--

SELECT pg_catalog.setval('upep.upep_molecular_types_id_seq', 1, false);


--
-- Name: upep_organisms_id_seq; Type: SEQUENCE SET; Schema: upep; Owner: postgres
--

SELECT pg_catalog.setval('upep.upep_organisms_id_seq', 1, false);


--
-- Name: upep_ref_seq_db_id_seq; Type: SEQUENCE SET; Schema: upep; Owner: postgres
--

SELECT pg_catalog.setval('upep.upep_ref_seq_db_id_seq', 1, false);


--
-- Name: upep_ref_seq_entries_id_seq; Type: SEQUENCE SET; Schema: upep; Owner: postgres
--

SELECT pg_catalog.setval('upep.upep_ref_seq_entries_id_seq', 1, false);


--
-- Name: upep_accessions upep_accessions_pkey; Type: CONSTRAINT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_accessions
    ADD CONSTRAINT upep_accessions_pkey PRIMARY KEY (id);


--
-- Name: upep_blast_db upep_blast_db_pkey; Type: CONSTRAINT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_blast_db
    ADD CONSTRAINT upep_blast_db_pkey PRIMARY KEY (id);


--
-- Name: upep_features upep_features_pkey; Type: CONSTRAINT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_features
    ADD CONSTRAINT upep_features_pkey PRIMARY KEY (id);


--
-- Name: upep_gene_identifiers upep_gis_id_pk; Type: CONSTRAINT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_gene_identifiers
    ADD CONSTRAINT upep_gis_id_pk PRIMARY KEY (id);


--
-- Name: upep_molecular_types upep_molecular_types_pkey; Type: CONSTRAINT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_molecular_types
    ADD CONSTRAINT upep_molecular_types_pkey PRIMARY KEY (id);


--
-- Name: upep_organisms upep_organisms_pkey; Type: CONSTRAINT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_organisms
    ADD CONSTRAINT upep_organisms_pkey PRIMARY KEY (id);


--
-- Name: upep_ref_seq_db upep_ref_seq_db_pkey; Type: CONSTRAINT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_ref_seq_db
    ADD CONSTRAINT upep_ref_seq_db_pkey PRIMARY KEY (id);


--
-- Name: upep_ref_seq_entries upep_ref_seq_entries_pkey; Type: CONSTRAINT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_ref_seq_entries
    ADD CONSTRAINT upep_ref_seq_entries_pkey PRIMARY KEY (id);


--
-- Name: upep_features upep_features_upep_ref_seq_entries_id_fk; Type: FK CONSTRAINT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_features
    ADD CONSTRAINT upep_features_upep_ref_seq_entries_id_fk FOREIGN KEY (ref_seq_entry_id) REFERENCES upep.upep_ref_seq_entries(id);


--
-- Name: upep_ref_seq_entries upep_ref_seq_entries_upep_accessions_id_fk; Type: FK CONSTRAINT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_ref_seq_entries
    ADD CONSTRAINT upep_ref_seq_entries_upep_accessions_id_fk FOREIGN KEY (accession_id) REFERENCES upep.upep_accessions(id);


--
-- Name: upep_ref_seq_entries upep_ref_seq_entries_upep_gis_id_fk; Type: FK CONSTRAINT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_ref_seq_entries
    ADD CONSTRAINT upep_ref_seq_entries_upep_gis_id_fk FOREIGN KEY (gi_id) REFERENCES upep.upep_gene_identifiers(id);


--
-- Name: upep_ref_seq_entries upep_ref_seq_entries_upep_molecular_types_id_fk; Type: FK CONSTRAINT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_ref_seq_entries
    ADD CONSTRAINT upep_ref_seq_entries_upep_molecular_types_id_fk FOREIGN KEY (molecular_type_id) REFERENCES upep.upep_molecular_types(id);


--
-- Name: upep_ref_seq_entries upep_ref_seq_entries_upep_organisms_id_fk; Type: FK CONSTRAINT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_ref_seq_entries
    ADD CONSTRAINT upep_ref_seq_entries_upep_organisms_id_fk FOREIGN KEY (organism_id) REFERENCES upep.upep_organisms(id);


--
-- Name: upep_ref_seq_entries upep_ref_seq_entries_upep_ref_seq_db_id_fk; Type: FK CONSTRAINT; Schema: upep; Owner: postgres
--

ALTER TABLE ONLY upep.upep_ref_seq_entries
    ADD CONSTRAINT upep_ref_seq_entries_upep_ref_seq_db_id_fk FOREIGN KEY (ref_seq_db_id) REFERENCES upep.upep_ref_seq_db(id);


--
-- PostgreSQL database dump complete
--


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop schema upep cascade;