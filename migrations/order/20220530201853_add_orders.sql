-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id         bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id    bigint NOT NULL,
    status     bigint NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
