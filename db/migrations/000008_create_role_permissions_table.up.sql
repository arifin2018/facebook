CREATE TABLE IF NOT EXISTS role_permisssions (
    role_id int NOT NULL,
    permissions_id int NOT NULL,
    CONSTRAINT role_permisssions_pk PRIMARY KEY (role_id, permissions_id),
    CONSTRAINT FK_role_id_role_permisssions FOREIGN KEY (role_id) REFERENCES roles (id),
    CONSTRAINT FK_permissions_id_role_permisssions FOREIGN KEY (permissions_id) REFERENCES permissions (id)
);