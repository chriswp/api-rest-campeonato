CREATE TABLE users
(
    id         UUID PRIMARY KEY    NOT NULL,
    name       VARCHAR(100)        NOT NULL,
    email      VARCHAR(100) UNIQUE NOT NULL,
    password   VARCHAR(60)         NOT NULL,
    created_at TIMESTAMP           NOT NULL DEFAULT now(),
    updated_at TIMESTAMP                    DEFAULT now()
);

CREATE TABLE football_fans
(
    id         UUID PRIMARY KEY    NOT NULL,
    name       VARCHAR(100)        NOT NULL,
    email      VARCHAR(100) UNIQUE NOT NULL,
    team       VARCHAR(60)         NOT NULL,
    created_at TIMESTAMP           NOT NULL DEFAULT now(),
    updated_at TIMESTAMP                    DEFAULT now()
);