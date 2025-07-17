/* SUPERADMIN */
INSERT INTO public.roles
(id, application_id, "name", description, "type", status, inactive_date, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at, resources)
VALUES(1, 1, 'SUPERUSER', 'Superadmin account', 'EXTERNAL', 'ACTIVE', NULL, 'SYSTEM', '2025-05-06 11:46:00.293', 'SYSTEM', '2025-05-06 11:46:00.293', NULL, NULL, NULL);

INSERT INTO public.users
(id, username, description, full_name, email, phone_number, "password", "type", status, start_date, end_date, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at, is_default_password, business_unit_level, business_unit_location_id, business_unit_location, business_unit_assignment_status, supplier_id, supplier_name, assigned_roles)
VALUES(1, 'admin1', 'Admin 1 Description', 'Admin 1', 'ehalimi@delinnce.com', '021123123', 'ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f', 'EXTERNAL', 'ACTIVE', NULL, NULL, 'SYSTEM', '2025-05-06 11:46:05.140', 'SYSTEM', '2025-05-06 11:46:05.140', NULL, NULL, true, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

INSERT INTO user_role_mappings(user_id, role_id, created_at, created_by, updated_at, updated_by) 
VALUES (1, 1, now(), 'SYSTEM', now(), 'SYSTEM');