CREATE TABLE resources (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    menu_id INT NOT NULL REFERENCES menus(id) ON DELETE CASCADE,
    submenu_id INT NOT NULL REFERENCES submenus(id) ON DELETE CASCADE,
    function_id INT NOT NULL REFERENCES functions(id) ON DELETE CASCADE,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_by VARCHAR NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_by VARCHAR NULL,
    deleted_at TIMESTAMPTZ NULL
);        

CREATE UNIQUE INDEX unique_resources_name_deleted 
    ON resources (name)
    WHERE deleted_at IS null and deleted_by is NULL;

CREATE UNIQUE INDEX unique_resources_name_menu_submenu_function_deleted 
    ON resources (name, menu_id, submenu_id, function_id)
    WHERE deleted_at IS null and deleted_by is NULL;