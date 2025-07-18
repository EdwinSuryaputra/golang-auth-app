CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    full_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    phone_number VARCHAR NULL,
    password VARCHAR NOT NULL,
    is_default_password BOOLEAN NOT NULL DEFAULT true,
    type VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    start_date TIMESTAMPTZ NULL,
    end_date TIMESTAMPTZ NULL,
    assigned_roles JSONB NULL,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_by VARCHAR NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_by VARCHAR NULL,
    deleted_at TIMESTAMPTZ NULL
);                                                                                                        

CREATE UNIQUE INDEX unique_users_username_deleted 
    ON users (username)
    WHERE deleted_at IS null and deleted_by is NULL;

CREATE UNIQUE INDEX unique_users_email_deleted 
    ON users (email)
    WHERE deleted_at IS null and deleted_by is NULL;