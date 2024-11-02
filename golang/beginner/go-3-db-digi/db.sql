-- public.accounts definition

-- Drop table

-- DROP TABLE public.accounts;

CREATE TABLE public.accounts (
	account_id int8 GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	name varchar NOT NULL,
	CONSTRAINT account_pk PRIMARY KEY (account_id)
);

-- public.auths definition

-- Drop table

-- DROP TABLE public.auths;

CREATE TABLE public.auths (
	auth_id int8 GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	account_id int8 NOT NULL,
	username varchar NOT NULL,
	"password" varchar NOT NULL,
	CONSTRAINT auths_pk PRIMARY KEY (auth_id),
	CONSTRAINT auths_unique UNIQUE (account_id),
	CONSTRAINT auths_unique_username UNIQUE (username)
);