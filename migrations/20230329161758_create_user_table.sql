-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  email varchar(100) not null,
  password varchar(128) not null,  
  name varchar(64) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
