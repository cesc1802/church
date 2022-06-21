CREATE TABLE IF NOT EXISTS user_groups (
                                           group_id INTEGER NOT NULL,
                                           user_id INTEGER NOT NULL ,
                                           created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,

                                           CONSTRAINT PK_GROUP_USER PRIMARY KEY (group_id, user_id)
    );