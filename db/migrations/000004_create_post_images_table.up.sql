CREATE TABLE IF NOT EXISTS post_images (
    ID int NOT NULL AUTO_INCREMENT,
    Post_id int NOT NULL,
    Url text NOT NULL,
    Created_at TIMESTAMP,
    Updated_at TIMESTAMP,
    Deleted_at TIMESTAMP,
    PRIMARY KEY (ID)
);