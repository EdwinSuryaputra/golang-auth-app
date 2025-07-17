CREATE TABLE user_role_mappings (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    role_id INT REFERENCES roles(id) ON DELETE CASCADE,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_by VARCHAR NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_by VARCHAR NULL,
    deleted_at TIMESTAMPTZ NULL
);                                                                                                        

CREATE UNIQUE INDEX unique_user_role_mappings_user_id_role_id_deleted 
    ON user_role_mappings (user_id, role_id)
    WHERE deleted_at IS null and deleted_by is NULL;