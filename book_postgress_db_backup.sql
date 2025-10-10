--
-- PostgreSQL database dump
--

--
-- TOC entry 222 (class 1255 OID 33726)
-- Name: fn_books_get_by_id(uuid); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.fn_books_get_by_id(p_id uuid) RETURNS TABLE(id uuid, name character varying, price real)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT id, name, price
    FROM books
    WHERE id = p_id;
END;
$$;


ALTER FUNCTION public.fn_books_get_by_id(p_id uuid) OWNER TO postgres;

--
-- TOC entry 220 (class 1255 OID 33724)
-- Name: fn_employee_get_by_id(uuid); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.fn_employee_get_by_id(p_id uuid) RETURNS TABLE(id uuid, name character varying, salary real)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT employee.id, employee.name, employee.salary
    FROM employee
    WHERE employee.id = p_id;
END;
$$;


ALTER FUNCTION public.fn_employee_get_by_id(p_id uuid) OWNER TO postgres;

--
-- TOC entry 221 (class 1255 OID 33725)
-- Name: sp_books_create(uuid, character varying, real); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.sp_books_create(p_id uuid, p_name character varying, p_price real) RETURNS uuid
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO books (id, name, price)
    VALUES (p_id, p_name, p_price);
	return p_id;
END;
$$;


ALTER FUNCTION public.sp_books_create(p_id uuid, p_name character varying, p_price real) OWNER TO postgres;

--
-- TOC entry 219 (class 1255 OID 33723)
-- Name: sp_employee_create(uuid, character varying, real); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.sp_employee_create(p_id uuid, p_name character varying, p_salary real) RETURNS uuid
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO employee (id, name, salary)
    VALUES (p_id, p_name, p_salary);
	return p_id;
END;
$$;


ALTER FUNCTION public.sp_employee_create(p_id uuid, p_name character varying, p_salary real) OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 217 (class 1259 OID 33707)
-- Name: books; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.books (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying NOT NULL,
    price real
);


ALTER TABLE public.books OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 33715)
-- Name: employee; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.employee (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying NOT NULL,
    salary real
);


ALTER TABLE public.employee OWNER TO postgres;

--
-- TOC entry 4900 (class 0 OID 33707)
-- Dependencies: 217
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 4901 (class 0 OID 33715)
-- Dependencies: 218
-- Data for Name: employee; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.employee VALUES ('1c302d52-d886-4a52-a1c8-a216b826c64c', 'john doe', 15000);


--
-- TOC entry 4752 (class 2606 OID 33714)
-- Name: books books_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);


--
-- TOC entry 4754 (class 2606 OID 33722)
-- Name: employee employee_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_pkey PRIMARY KEY (id);


-- Completed on 2025-10-10 16:52:11

--
-- PostgreSQL database dump complete
--

