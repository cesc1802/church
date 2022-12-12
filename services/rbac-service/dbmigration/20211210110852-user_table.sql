
CREATE TABLE IF NOT EXISTS users
(
    id
    BIGINT
    PRIMARY
    KEY,
    UserID
    VARCHAR
(
    50
),
    PASSWORD VARCHAR
(
    50
)
    );
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
    150
) NOT NULL,
    RedirectURIs VARCHAR (150) ARRAY,
    GrantTypes VARCHAR (150) ARRAY,
    ResponseTypes VARCHAR (150) ARRAY,
    Scopes VARCHAR (150) ARRAY,
    Audience VARCHAR (150) ARRAY,
    Public boolean
    );