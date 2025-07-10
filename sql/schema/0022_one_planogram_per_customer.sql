-- +goose Up
-- Enforce that only one planogram can be assigned to a customer location
ALTER TABLE planogram_customer_locations
ADD CONSTRAINT unique_customer_location_per_planogram UNIQUE (customer_location_id);

-- +goose Down
ALTER TABLE planogram_customer_locations
DROP CONSTRAINT unique_customer_location_per_planogram;
