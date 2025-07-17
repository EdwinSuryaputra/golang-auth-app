CREATE TABLE temp_roles (
    id SERIAL PRIMARY KEY,
    role_id INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    type VARCHAR NOT NULL,
    inactive_date TIMESTAMPTZ NULL,
    resources JSONB NULL,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_by VARCHAR NULL,
    deleted_at TIMESTAMPTZ NULL
);                                                                                                             

CREATE UNIQUE INDEX unique_temp_roles_deleted 
    ON temp_roles (role_id)
    WHERE deleted_at IS null and deleted_by is NULL;