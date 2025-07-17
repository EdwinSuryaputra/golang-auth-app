CREATE TABLE applications (
    id SERIAL PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL,
    public_name VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_by VARCHAR NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_by VARCHAR NULL,
    deleted_at TIMESTAMPTZ NULL
);                                                                                                        

CREATE UNIQUE INDEX unique_applications_name_deleted 
    ON applications (name)
    WHERE deleted_at IS null and deleted_by is NULL;
