CREATE TABLE IF NOT EXISTS permissions (
    ID int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    Created_at TIMESTAMP,
    Updated_at TIMESTAMP,
    Deleted_at TIMESTAMP,
    PRIMARY KEY (ID),
    UNIQUE (name)
);