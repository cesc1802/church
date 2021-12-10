
-- +migrate Up
CREATE TABLE IF NOT EXISTS group_roles (
    group_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL ,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT PK_GROUP_ROLE PRIMARY KEY (group_id, role_id)
);
-- +migrate Down
DROP TABLE IF EXISTS group_roles;