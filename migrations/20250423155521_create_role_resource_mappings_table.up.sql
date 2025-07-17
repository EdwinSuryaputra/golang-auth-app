CREATE TABLE role_resource_mappings (
    id SERIAL PRIMARY KEY,
    role_id INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    resource_id INT NOT NULL REFERENCES resources(id) ON DELETE CASCADE,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_by VARCHAR NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_by VARCHAR NULL,
    deleted_at TIMESTAMPTZ NULL
);                                                                                                         

CREATE UNIQUE INDEX unique_role_resource_mappings_deleted 
    ON role_resource_mappings (role_id, resource_id)
    WHERE deleted_at IS null and deleted_by is NULL;