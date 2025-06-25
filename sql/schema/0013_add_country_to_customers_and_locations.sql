-- +goose Up
ALTER TABLE customers
ADD COLUMN country CHAR(3) NOT NULL DEFAULT 'USA';


ALTER TABLE customer_locations
ADD COLUMN country CHAR(3) NOT NULL DEFAULT 'USA';

-- +goose Down
ALTER TABLE customers
DROP COLUMN country; 

ALTER TABLE customer_locations
DROP COLUMN country;