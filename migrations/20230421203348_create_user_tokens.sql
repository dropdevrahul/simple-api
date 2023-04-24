-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE user_tokens(
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  token varchar(100) not null unique,
  user_id uuid REFERENCES users
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_tokens;
-- +goose StatementEnd
