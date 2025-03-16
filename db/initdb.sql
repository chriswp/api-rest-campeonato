CREATE DATABASE "football";

\c football

CREATE TABLE public.users
(
    id         UUID PRIMARY KEY    NOT NULL,
    name       VARCHAR(100)        NOT NULL,
    email      VARCHAR(100) UNIQUE NOT NULL,
    password   VARCHAR(60)         NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

INSERT INTO public.users (id, name, email, password, created_at, updated_at) VALUES ('b8a6fbc2-3d7f-45a6-9c5d-0c8e4a9d72f3', 'user', 'user@test.com', '$2b$12$Am6A9CfyOgIXHoCdj3PjQu5IEHeYEfSk9Bnl1uYm91/Cb1jg6puX2', '2025-03-10 00:09:16.000000', '2025-03-10 00:09:16.000000');

CREATE TABLE football_fans
(
    id         UUID PRIMARY KEY    NOT NULL,
    name       VARCHAR(100)        NOT NULL,
    email      VARCHAR(100) UNIQUE NOT NULL,
    team       VARCHAR(60)         NOT NULL,
    created_at TIMESTAMP           NOT NULL DEFAULT now(),
    updated_at TIMESTAMP                    DEFAULT now()
);