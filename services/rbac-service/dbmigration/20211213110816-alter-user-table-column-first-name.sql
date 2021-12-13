
-- +migrate Up
ALTER TABLE users ADD COLUMN first_name VARCHAR (50);
-- +migrate Down
ALTER TABLE users DROP COLUMN first_name;