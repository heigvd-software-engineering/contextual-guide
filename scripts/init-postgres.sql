--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.6
-- Dumped by pg_dump version 10.4 (Debian 10.4-2.pgdg90+1)

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
-- Name: gotrue; Type: SCHEMA; Schema: -; Owner: gotrue
--

CREATE DATABASE gotrue;

\connect gotrue;

CREATE SCHEMA gotrue;

CREATE USER gotrue WITH PASSWORD 'gotrue';

ALTER SCHEMA gotrue OWNER TO gotrue;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: audit_log_entries; Type: TABLE; Schema: gotrue; Owner: gotrue
--

CREATE TABLE gotrue.audit_log_entries (
    instance_id character varying(255),
    id character varying(255) NOT NULL,
    payload json,
    created_at timestamp with time zone
);


ALTER TABLE gotrue.audit_log_entries OWNER TO gotrue;

--
-- Name: instances; Type: TABLE; Schema: gotrue; Owner: gotrue
--

CREATE TABLE gotrue.instances (
    id character varying(255) NOT NULL,
    uuid character varying(255),
    raw_base_config text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE gotrue.instances OWNER TO gotrue;

--
-- Name: refresh_tokens; Type: TABLE; Schema: gotrue; Owner: gotrue
--

CREATE TABLE gotrue.refresh_tokens (
    instance_id character varying(255),
    id bigint NOT NULL,
    token character varying(255),
    user_id character varying(255),
    revoked boolean,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE gotrue.refresh_tokens OWNER TO gotrue;

--
-- Name: refresh_tokens_id_seq; Type: SEQUENCE; Schema: gotrue; Owner: gotrue
--

CREATE SEQUENCE gotrue.refresh_tokens_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gotrue.refresh_tokens_id_seq OWNER TO gotrue;

--
-- Name: refresh_tokens_id_seq; Type: SEQUENCE OWNED BY; Schema: gotrue; Owner: gotrue
--

ALTER SEQUENCE gotrue.refresh_tokens_id_seq OWNED BY gotrue.refresh_tokens.id;


--
-- Name: schema_migration; Type: TABLE; Schema: gotrue; Owner: gotrue
--

CREATE TABLE gotrue.schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE gotrue.schema_migration OWNER TO gotrue;

--
-- Name: users; Type: TABLE; Schema: gotrue; Owner: gotrue
--

CREATE TABLE gotrue.users (
    instance_id character varying(255),
    id character varying(255) NOT NULL,
    aud character varying(255),
    role character varying(255),
    email character varying(255),
    encrypted_password character varying(255),
    confirmed_at timestamp with time zone,
    invited_at timestamp with time zone,
    confirmation_token character varying(255),
    confirmation_sent_at timestamp with time zone,
    recovery_token character varying(255),
    recovery_sent_at timestamp with time zone,
    email_change_token character varying(255),
    email_change character varying(255),
    email_change_sent_at timestamp with time zone,
    last_sign_in_at timestamp with time zone,
    raw_app_meta_data json,
    raw_user_meta_data json,
    is_super_admin boolean,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE gotrue.users OWNER TO gotrue;

--
-- Name: refresh_tokens id; Type: DEFAULT; Schema: gotrue; Owner: gotrue
--

ALTER TABLE ONLY gotrue.refresh_tokens ALTER COLUMN id SET DEFAULT nextval('gotrue.refresh_tokens_id_seq'::regclass);


--
-- Data for Name: audit_log_entries; Type: TABLE DATA; Schema: gotrue; Owner: gotrue
--

COPY gotrue.audit_log_entries (instance_id, id, payload, created_at) FROM stdin;
\.


--
-- Data for Name: instances; Type: TABLE DATA; Schema: gotrue; Owner: gotrue
--

COPY gotrue.instances (id, uuid, raw_base_config, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: refresh_tokens; Type: TABLE DATA; Schema: gotrue; Owner: gotrue
--

COPY gotrue.refresh_tokens (instance_id, id, token, user_id, revoked, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: schema_migration; Type: TABLE DATA; Schema: gotrue; Owner: gotrue
--

COPY gotrue.schema_migration (version) FROM stdin;
20171026211738
20171026211808
20171026211834
20180103212743
20180108183307
20180119214651
20180125194653
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: gotrue; Owner: gotrue
--

COPY gotrue.users (instance_id, id, aud, role, email, encrypted_password, confirmed_at, invited_at, confirmation_token, confirmation_sent_at, recovery_token, recovery_sent_at, email_change_token, email_change, email_change_sent_at, last_sign_in_at, raw_app_meta_data, raw_user_meta_data, is_super_admin, created_at, updated_at) FROM stdin;
\.


--
-- Name: refresh_tokens_id_seq; Type: SEQUENCE SET; Schema: gotrue; Owner: gotrue
--

SELECT pg_catalog.setval('gotrue.refresh_tokens_id_seq', 1, true);


--
-- Name: audit_log_entries idx_16496_primary; Type: CONSTRAINT; Schema: gotrue; Owner: gotrue
--

ALTER TABLE ONLY gotrue.audit_log_entries
    ADD CONSTRAINT idx_16496_primary PRIMARY KEY (id);


--
-- Name: instances idx_16502_primary; Type: CONSTRAINT; Schema: gotrue; Owner: gotrue
--

ALTER TABLE ONLY gotrue.instances
    ADD CONSTRAINT idx_16502_primary PRIMARY KEY (id);


--
-- Name: refresh_tokens idx_16510_primary; Type: CONSTRAINT; Schema: gotrue; Owner: gotrue
--

ALTER TABLE ONLY gotrue.refresh_tokens
    ADD CONSTRAINT idx_16510_primary PRIMARY KEY (id);


--
-- Name: users idx_16520_primary; Type: CONSTRAINT; Schema: gotrue; Owner: gotrue
--

ALTER TABLE ONLY gotrue.users
    ADD CONSTRAINT idx_16520_primary PRIMARY KEY (id);


--
-- Name: idx_16496_audit_logs_instance_id_idx; Type: INDEX; Schema: gotrue; Owner: gotrue
--

CREATE INDEX idx_16496_audit_logs_instance_id_idx ON gotrue.audit_log_entries USING btree (instance_id);


--
-- Name: idx_16510_refresh_tokens_instance_id_idx; Type: INDEX; Schema: gotrue; Owner: gotrue
--

CREATE INDEX idx_16510_refresh_tokens_instance_id_idx ON gotrue.refresh_tokens USING btree (instance_id);


--
-- Name: idx_16510_refresh_tokens_instance_id_user_id_idx; Type: INDEX; Schema: gotrue; Owner: gotrue
--

CREATE INDEX idx_16510_refresh_tokens_instance_id_user_id_idx ON gotrue.refresh_tokens USING btree (instance_id, user_id);


--
-- Name: idx_16510_refresh_tokens_token_idx; Type: INDEX; Schema: gotrue; Owner: gotrue
--

CREATE INDEX idx_16510_refresh_tokens_token_idx ON gotrue.refresh_tokens USING btree (token);


--
-- Name: idx_16517_version_idx; Type: INDEX; Schema: gotrue; Owner: gotrue
--

CREATE UNIQUE INDEX idx_16517_version_idx ON gotrue.schema_migration USING btree (version);


--
-- Name: idx_16520_users_instance_id_email_idx; Type: INDEX; Schema: gotrue; Owner: gotrue
--

CREATE INDEX idx_16520_users_instance_id_email_idx ON gotrue.users USING btree (instance_id, email);


--
-- Name: idx_16520_users_instance_id_idx; Type: INDEX; Schema: gotrue; Owner: gotrue
--

CREATE INDEX idx_16520_users_instance_id_idx ON gotrue.users USING btree (instance_id);


--
-- PostgreSQL database dump complete
--

