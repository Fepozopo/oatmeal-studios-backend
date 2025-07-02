-- +goose Up
ALTER TABLE customers ADD COLUMN free_shipping BOOLEAN NOT NULL DEFAULT FALSE;

-- +goose Down
ALTER TABLE customers DROP COLUMN free_shipping;
