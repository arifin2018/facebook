CREATE TABLE IF NOT EXISTS user_roles (
    user_id int NOT NULL,
    role_id int NOT NULL,
    CONSTRAINT user_roles_pk PRIMARY KEY (user_id, role_id),
    CONSTRAINT FK_role_id_user_roles FOREIGN KEY (role_id) REFERENCES roles (id),
    CONSTRAINT FK_user_id_user_roles FOREIGN KEY (user_id) REFERENCES users (id)
);