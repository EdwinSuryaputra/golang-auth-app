CREATE TABLE bu_request_buckets (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    business_unit_level VARCHAR NOT NULL,
    business_unit_location_id INT NOT NULL,
    business_unit_location VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    reviewed_by VARCHAR NULL,
    reviewed_at TIMESTAMPTZ NULL,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_by VARCHAR NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_by VARCHAR NULL,
    deleted_at TIMESTAMPTZ NULL
);                                                                                                        