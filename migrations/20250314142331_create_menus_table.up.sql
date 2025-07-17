CREATE TABLE menus (
    id SERIAL PRIMARY KEY,
    application_id INT NOT NULL REFERENCES applications(id) ON DELETE CASCADE,
    name VARCHAR NOT NULL,
    public_name VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_by VARCHAR NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_by VARCHAR NULL,
    deleted_at TIMESTAMPTZ NULL
);                                                                                                             

CREATE UNIQUE INDEX unique_menus_application_id_name_deleted 
    ON menus (application_id, name)
    WHERE deleted_at IS null and deleted_by is NULL;