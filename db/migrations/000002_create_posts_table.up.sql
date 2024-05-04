CREATE TABLE IF NOT EXISTS posts (
    Id int NOT NULL AUTO_INCREMENT,
    User_id int NOT NULL,
    Content text NOT NULL,
    Created_at TIMESTAMP,
    Updated_at TIMESTAMP,
    Deleted_at TIMESTAMP,
    PRIMARY KEY (Id)
);