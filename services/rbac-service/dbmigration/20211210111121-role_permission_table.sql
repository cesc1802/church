
-- +migrate Up
CREATE TABLE IF NOT EXISTS role_permissions (
    role_id INTEGER NOT NULL ,
    permission_id INTEGER NOT NULL ,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ,

    CONSTRAINT PK_ROLE_PERMISSION PRIMARY KEY (role_id, permission_id)
);
-- +migrate Down
DROP TABLE IF EXISTS role_permissions;