-- +goose Up
-- +goose StatementBegin
CREATE TABLE cart_items (
     user_id          bigint NOT NULL,
     product_id  bigint NOT NULL,
     count          bigint NOT NULL,
     UNIQUE(user_id, product_id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE subscribers;
-- +goose StatementEnd
