-- +migrate Up
CREATE TABLE IF NOT EXISTS users
(
    id
    SERIAL
    NOT
    NULL
    PRIMARY
    KEY,
    login_id
    VARCHAR
(
    50
),
    fb_id VARCHAR
(
    50
),
    gg_id VARCHAR
(
    50
),
    password VARCHAR
(
    255
) NOT NULL,
    status INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
    );
-- +migrate Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS clients;

CREATE TABLE IF NOT EXISTS users
(
    id
    SERIAL
    NOT
    NULL
    PRIMARY
    KEY,
    username
    VARCHAR
(
    50
) NOT NULL UNIQUE,
    password VARCHAR
(
    50
)
    );

CREATE TABLE IF NOT EXISTS clients
(
    id VARCHAR (50)
    NOT
    NULL
    PRIMARY
    KEY,
    Secret
    VARCHAR
(
    50
) NOT NULL,
    RedirectURIs text[],
    GrantTypes text[],
    ResponseTypes text[],
    Scopes text[],
    Audience text[],
    Public boolean
    );