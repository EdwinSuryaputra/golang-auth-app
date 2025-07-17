CREATE TABLE submenus (
    id SERIAL PRIMARY KEY,
    menu_id INT NOT NULL REFERENCES menus(id) ON DELETE CASCADE,
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

CREATE UNIQUE INDEX unique_submenus_menu_id_name_deleted 
    ON submenus (menu_id, name)
    WHERE deleted_at IS null and deleted_by is NULL;