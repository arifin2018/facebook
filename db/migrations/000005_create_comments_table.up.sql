CREATE TABLE IF NOT EXISTS comments (
    ID int NOT NULL AUTO_INCREMENT,
    User_id int NOT NULL,
    Post_id int NOT NULL,
    Content text NOT NULL,
    Created_at TIMESTAMP,
    Updated_at TIMESTAMP,
    Deleted_at TIMESTAMP,
    PRIMARY KEY (ID),
    FOREIGN KEY (Post_id)
      REFERENCES posts (id)
);