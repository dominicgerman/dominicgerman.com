-- +goose Up
CREATE TABLE posts (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(255),
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    updated DATETIME NOT NULL
);

-- +goose Down
DROP TABLE posts;