
-- +migrate Up
CREATE TABLE IF NOT EXISTS user_roles (
    user_id INTEGER NOT NULL ,
    role_id INTEGER NOT NULL ,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT PK_USER_ROLE PRIMARY KEY (user_id, role_id)
);
-- +migrate Down
DROP TABLE IF EXISTS user_roles;