CREATE TABLE activity_logs (
    id SERIAL PRIMARY KEY,
    activity_log_id VARCHAR NOT NULL,
    message VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_by VARCHAR NULL,
    updated_at TIMESTAMPTZ NULL,
    deleted_by VARCHAR NULL,
    deleted_at TIMESTAMPTZ null
);