-- +goose Up
-- +goose StatementBegin
CREATE TABLE reserves (
    order_id   bigint NOT NULL,
    product_id bigint NOT NULL,
    count      bigint NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reserves;
-- +goose StatementEnd
