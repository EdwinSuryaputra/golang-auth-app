ALTER TABLE roles 
DROP COLUMN activity_log_id;

ALTER TABLE users 
DROP COLUMN activity_log_id;

ALTER TABLE bu_request_buckets 
DROP COLUMN activity_log_id;

DROP UNIQUE INDEX unique_role_activity_log_id_deleted;

DROP UNIQUE INDEX unique_user_activity_log_id_deleted;

DROP UNIQUE INDEX unique_burb_activity_log_id_deleted;