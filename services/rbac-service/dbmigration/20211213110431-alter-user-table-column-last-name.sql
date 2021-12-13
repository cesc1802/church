
-- +migrate Up
ALTER TABLE users ADD COLUMN last_name VARCHAR (50);
-- +migrate Down
ALTER TABLE users DROP COLUMN last_name;