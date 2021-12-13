
-- +migrate Up
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL NOT NULL PRIMARY KEY ,
    name VARCHAR (50) NOT NULL ,
    status INTEGER DEFAULT 1,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);
-- +migrate Down
DROP TABLE IF EXISTS roles;