-- +goose Up
ALTER TABLE customer_locations
    ADD COLUMN business_name TEXT NOT NULL DEFAULT '',
    ADD COLUMN contact_name TEXT;

-- +goose Down
ALTER TABLE customer_locations
    DROP COLUMN contact_name,
    DROP COLUMN business_name;
