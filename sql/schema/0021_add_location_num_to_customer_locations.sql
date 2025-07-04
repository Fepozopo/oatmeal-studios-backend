-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE customer_locations ADD COLUMN location_num INTEGER;

-- +goose Down
-- SQL in section 'Down' is executed when this migration is rolled back
ALTER TABLE customer_locations DROP COLUMN location_num;
