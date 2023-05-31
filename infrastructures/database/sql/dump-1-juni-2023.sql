--
-- PostgreSQL database dump
--

-- Dumped from database version 11.19 (Ubuntu 11.19-1.pgdg20.04+1)
-- Dumped by pg_dump version 14.6

-- Started on 2023-06-01 04:36:36 WIB

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

DROP DATABASE dwbejsql;
--
-- TOC entry 4043 (class 1262 OID 11047367)
-- Name: dwbejsql; Type: DATABASE; Schema: -; Owner: dwbejsql
--

CREATE DATABASE dwbejsql WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.UTF-8';


ALTER DATABASE dwbejsql OWNER TO dwbejsql;

\connect dwbejsql

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
-- TOC entry 24 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 4045 (class 0 OID 0)
-- Dependencies: 24
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

--
-- TOC entry 226 (class 1259 OID 12447044)
-- Name: device_mode; Type: TABLE; Schema: public; Owner: dwbejsql
--

CREATE TABLE public.device_mode (
    mode_id integer NOT NULL,
    mode_name character varying,
    date_created timestamp without time zone
);


ALTER TABLE public.device_mode OWNER TO dwbejsql;

--
-- TOC entry 222 (class 1259 OID 11047368)
-- Name: device_status; Type: TABLE; Schema: public; Owner: dwbejsql
--

CREATE TABLE public.device_status (
    status_id integer NOT NULL,
    status_name character varying,
    date_created timestamp without time zone
);


ALTER TABLE public.device_status OWNER TO dwbejsql;

--
-- TOC entry 223 (class 1259 OID 11047374)
-- Name: devices; Type: TABLE; Schema: public; Owner: dwbejsql
--

CREATE TABLE public.devices (
    device_id uuid NOT NULL,
    antares_id character varying,
    device_name character varying,
    device_location character varying,
    device_label character varying,
    status_id integer,
    brand character varying,
    user_id uuid,
    latitude character varying,
    longitude character varying,
    date_created timestamp without time zone,
    date_updated timestamp without time zone,
    mode_id integer
);


ALTER TABLE public.devices OWNER TO dwbejsql;

--
-- TOC entry 224 (class 1259 OID 11047380)
-- Name: roles; Type: TABLE; Schema: public; Owner: dwbejsql
--

CREATE TABLE public.roles (
    role_id integer NOT NULL,
    role_name character varying,
    date_created character varying
);


ALTER TABLE public.roles OWNER TO dwbejsql;

--
-- TOC entry 225 (class 1259 OID 11047386)
-- Name: users; Type: TABLE; Schema: public; Owner: dwbejsql
--

CREATE TABLE public.users (
    user_id uuid NOT NULL,
    role_id integer,
    full_name character varying,
    email character varying,
    password character varying,
    address character varying,
    avatar character varying,
    date_created timestamp without time zone,
    date_updated timestamp without time zone
);


ALTER TABLE public.users OWNER TO dwbejsql;

--
-- TOC entry 4037 (class 0 OID 12447044)
-- Dependencies: 226
-- Data for Name: device_mode; Type: TABLE DATA; Schema: public; Owner: dwbejsql
--

INSERT INTO public.device_mode VALUES (1, 'Otomatis', '2023-06-01 03:07:18.071393');
INSERT INTO public.device_mode VALUES (2, 'Manual', '2023-06-01 03:07:27.461798');


--
-- TOC entry 4033 (class 0 OID 11047368)
-- Dependencies: 222
-- Data for Name: device_status; Type: TABLE DATA; Schema: public; Owner: dwbejsql
--

INSERT INTO public.device_status VALUES (11, 'Aktif', '2023-03-30 06:14:05.677865');
INSERT INTO public.device_status VALUES (10, 'Tidak Aktif', '2023-03-30 06:14:21.438698');


--
-- TOC entry 4034 (class 0 OID 11047374)
-- Dependencies: 223
-- Data for Name: devices; Type: TABLE DATA; Schema: public; Owner: dwbejsql
--

INSERT INTO public.devices VALUES ('e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 'ff-ww-antares', 'Aerator Satu', 'Aquarium Kiri Hitam', 'Dummy', 11, 'Amara', 'a962321c-6b3a-4b92-8a70-9729a1f15b75', NULL, NULL, '2023-03-30 06:15:57.118381', '2023-06-01 04:32:51.180575', 2);


--
-- TOC entry 4035 (class 0 OID 11047380)
-- Dependencies: 224
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: dwbejsql
--

INSERT INTO public.roles VALUES (1, 'Super Admin', 'now()');


--
-- TOC entry 4036 (class 0 OID 11047386)
-- Dependencies: 225
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: dwbejsql
--

INSERT INTO public.users VALUES ('a962321c-6b3a-4b92-8a70-9729a1f15b75', 1, 'GuppyTech Admin', 'admin@guppytech.id', '$2a$12$xOeTPV2cIAcadxrGrkuxYemySlKNYoVjvtcAxvL1IEqY5Jk.XETb6
', 'Bandung', NULL, '2023-03-30 05:46:05.620484', '2023-03-30 05:46:05.620484');


--
-- TOC entry 3906 (class 2606 OID 12447048)
-- Name: device_mode device_mode_pk; Type: CONSTRAINT; Schema: public; Owner: dwbejsql
--

ALTER TABLE ONLY public.device_mode
    ADD CONSTRAINT device_mode_pk PRIMARY KEY (mode_id);


--
-- TOC entry 3898 (class 2606 OID 11047393)
-- Name: device_status device_status_pk; Type: CONSTRAINT; Schema: public; Owner: dwbejsql
--

ALTER TABLE ONLY public.device_status
    ADD CONSTRAINT device_status_pk PRIMARY KEY (status_id);


--
-- TOC entry 3900 (class 2606 OID 11047395)
-- Name: devices devices_pk; Type: CONSTRAINT; Schema: public; Owner: dwbejsql
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT devices_pk PRIMARY KEY (device_id);


--
-- TOC entry 3904 (class 2606 OID 11047397)
-- Name: users newtable_pk; Type: CONSTRAINT; Schema: public; Owner: dwbejsql
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT newtable_pk PRIMARY KEY (user_id);


--
-- TOC entry 3902 (class 2606 OID 11047399)
-- Name: roles role_pk; Type: CONSTRAINT; Schema: public; Owner: dwbejsql
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT role_pk PRIMARY KEY (role_id);


--
-- TOC entry 3907 (class 2606 OID 11047400)
-- Name: devices devices_fk; Type: FK CONSTRAINT; Schema: public; Owner: dwbejsql
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT devices_fk FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- TOC entry 3909 (class 2606 OID 12447052)
-- Name: devices mode_fk; Type: FK CONSTRAINT; Schema: public; Owner: dwbejsql
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT mode_fk FOREIGN KEY (mode_id) REFERENCES public.device_mode(mode_id);


--
-- TOC entry 3908 (class 2606 OID 11047405)
-- Name: devices status_fk; Type: FK CONSTRAINT; Schema: public; Owner: dwbejsql
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT status_fk FOREIGN KEY (status_id) REFERENCES public.device_status(status_id);


--
-- TOC entry 3910 (class 2606 OID 11047410)
-- Name: users users_fk; Type: FK CONSTRAINT; Schema: public; Owner: dwbejsql
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_fk FOREIGN KEY (role_id) REFERENCES public.roles(role_id);


--
-- TOC entry 4044 (class 0 OID 0)
-- Dependencies: 4043
-- Name: DATABASE dwbejsql; Type: ACL; Schema: -; Owner: dwbejsql
--

REVOKE CONNECT,TEMPORARY ON DATABASE dwbejsql FROM PUBLIC;


-- Completed on 2023-06-01 04:36:45 WIB

--
-- PostgreSQL database dump complete
--

