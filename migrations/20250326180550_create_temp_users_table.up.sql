CREATE TABLE temp_users (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    username VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    full_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    phone_number VARCHAR NULL,
    supplier_name VARCHAR NULL,
    assigned_roles JSONB NULL,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_by VARCHAR NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_by VARCHAR NULL,
    deleted_at TIMESTAMPTZ NULL
);                                                                                                        

CREATE UNIQUE INDEX unique_temp_users_deleted 
    ON temp_users (user_id)
    WHERE deleted_at IS null and deleted_by is NULL;