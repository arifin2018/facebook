CREATE TABLE IF NOT EXISTS users (
    Id int NOT NULL AUTO_INCREMENT,
    Name varchar(255) NOT NULL,
    Email varchar(255) NOT NULL,
    Password varchar(255) NOT NULL,
    Image varchar(255),
    PRIMARY KEY (Id),
    UNIQUE (Email)
);