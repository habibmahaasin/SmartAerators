--
-- PostgreSQL database dump
--

-- Dumped from database version 14.6
-- Dumped by pg_dump version 14.6

-- Started on 2023-03-30 06:18:26 WIB

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

DROP DATABASE smart_aerators;
--
-- TOC entry 3600 (class 1262 OID 17271)
-- Name: smart_aerators; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE smart_aerators WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'C';


ALTER DATABASE smart_aerators OWNER TO postgres;

\connect smart_aerators

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
-- TOC entry 3 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 3601 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 212 (class 1259 OID 17363)
-- Name: device_status; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.device_status (
    status_id integer NOT NULL,
    status_name character varying,
    date_created timestamp without time zone
);


ALTER TABLE public.device_status OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 17313)
-- Name: devices; Type: TABLE; Schema: public; Owner: postgres
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
    date_updated timestamp without time zone
);


ALTER TABLE public.devices OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 17279)
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    role_id integer NOT NULL,
    role_name character varying,
    date_created character varying
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 17272)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
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


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 3594 (class 0 OID 17363)
-- Dependencies: 212
-- Data for Name: device_status; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.device_status VALUES (11, 'Aktif', '2023-03-30 06:14:05.677865');
INSERT INTO public.device_status VALUES (10, 'Tidak Aktif', '2023-03-30 06:14:21.438698');


--
-- TOC entry 3593 (class 0 OID 17313)
-- Dependencies: 211
-- Data for Name: devices; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.devices VALUES ('e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 'ff-ww-antares', 'Aerator Satu', 'Aquarium Kiri Hitam', 'Dummy', 11, 'Amara', 'a962321c-6b3a-4b92-8a70-9729a1f15b75', NULL, NULL, '2023-03-30 06:15:57.118381', '2023-03-30 06:15:57.118381');


--
-- TOC entry 3592 (class 0 OID 17279)
-- Dependencies: 210
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.roles VALUES (1, 'Super Admin', 'now()');


--
-- TOC entry 3591 (class 0 OID 17272)
-- Dependencies: 209
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES ('a962321c-6b3a-4b92-8a70-9729a1f15b75', 1, 'GuppyTech Admin', 'admin@guppytech.id', '$2a$12$Ty7zvxYXKblB/GmCsNfRN./GbIgX.0Prdzo5/8k0Xt4czAcwbgF1e
', 'Bandung', NULL, '2023-03-30 05:46:05.620484', '2023-03-30 05:46:05.620484');


--
-- TOC entry 3448 (class 2606 OID 17369)
-- Name: device_status device_status_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_status
    ADD CONSTRAINT device_status_pk PRIMARY KEY (status_id);


--
-- TOC entry 3446 (class 2606 OID 17332)
-- Name: devices devices_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT devices_pk PRIMARY KEY (device_id);


--
-- TOC entry 3442 (class 2606 OID 17344)
-- Name: users newtable_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT newtable_pk PRIMARY KEY (user_id);


--
-- TOC entry 3444 (class 2606 OID 17297)
-- Name: roles role_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT role_pk PRIMARY KEY (role_id);


--
-- TOC entry 3450 (class 2606 OID 17357)
-- Name: devices devices_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT devices_fk FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- TOC entry 3451 (class 2606 OID 17387)
-- Name: devices status_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT status_fk FOREIGN KEY (status_id) REFERENCES public.device_status(status_id);


--
-- TOC entry 3449 (class 2606 OID 17308)
-- Name: users users_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_fk FOREIGN KEY (role_id) REFERENCES public.roles(role_id);


-- Completed on 2023-03-30 06:18:26 WIB

--
-- PostgreSQL database dump complete
--

