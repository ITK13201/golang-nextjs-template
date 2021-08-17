-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
   id INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
   name VARCHAR(128) NOT NULL UNIQUE,
   password VARCHAR(128) NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
