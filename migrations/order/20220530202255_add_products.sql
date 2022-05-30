-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
    order_id   bigint REFERENCES orders (id) ON DELETE CASCADE NOT NULL,
    product_id bigint NOT NULL,
    count      bigint NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd
