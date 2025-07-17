ALTER TABLE roles 
ADD COLUMN activity_log_id VARCHAR NULL;

ALTER TABLE users 
ADD COLUMN activity_log_id VARCHAR NULL;

ALTER TABLE bu_request_buckets 
ADD COLUMN activity_log_id VARCHAR NULL;

CREATE UNIQUE INDEX unique_role_activity_log_id_deleted 
    ON roles (activity_log_id)
    WHERE deleted_at IS null and deleted_by is NULL;

CREATE UNIQUE INDEX unique_user_qc_activity_log_id_deleted 
    ON users (activity_log_id)
    WHERE deleted_at IS null and deleted_by is NULL;

CREATE UNIQUE INDEX unique_burb_qc_activity_log_id_deleted 
    ON bu_request_buckets (activity_log_id)
    WHERE deleted_at IS null and deleted_by is NULL;

