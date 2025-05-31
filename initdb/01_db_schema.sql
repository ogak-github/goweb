--
-- PostgreSQL database setup for source database with logical replication
-- This is the source database that will send data to PowerSync
--

-- Enable logical replication
ALTER SYSTEM SET wal_level = 'logical';

-- Create a role/user with replication privileges for PowerSync
CREATE ROLE powersync_user WITH REPLICATION BYPASSRLS LOGIN PASSWORD 'my_secure_password';


-- Set up permissions for the PowerSync user
GRANT SELECT ON ALL TABLES IN SCHEMA public TO powersync_user;

-- To grant permissions to new created tables
ALTER DEFAULT PRIVILEGES IN SCHEMA public
  GRANT SELECT ON TABLES TO powersync_user;

-- Make sure sequences are also granted for user powersync_user
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO powersync_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT USAGE, SELECT ON SEQUENCES TO powersync_user;


-- Create publication for logical replication
CREATE PUBLICATION powersync FOR ALL TABLES;

-- Setup Database
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



--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: todo; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.todo (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    title character varying(255),
    content text,
    created_at timestamp without time zone,
    modify_at timestamp without time zone,
    user_id uuid
);


ALTER TABLE public.todo OWNER TO admin;

--
-- Name: users; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.users (
    username character varying(100) NOT NULL,
    password character varying(255) NOT NULL,
    full_name character varying(255),
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    email character varying(255) NOT NULL,
    CONSTRAINT username_lower_alnum_only CHECK (((username)::text ~ '^[a-z0-9]+$'::text)),
    CONSTRAINT valid_email_format CHECK (((email)::text ~ '^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'::text))
);


ALTER TABLE public.users OWNER TO admin;

--
-- Name: todo todo_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.todo
    ADD CONSTRAINT todo_pkey PRIMARY KEY (id);


--
-- Name: users unique_username; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT unique_username UNIQUE (username);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: todo todo_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.todo
    ADD CONSTRAINT todo_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--



