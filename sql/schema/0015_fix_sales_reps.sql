-- +goose Up
-- Drop sales_rep column from customers table
ALTER TABLE customers
    DROP COLUMN sales_rep;

-- Add sales_rep column to customer_locations table, referencing sales_reps(rep_code)
ALTER TABLE customer_locations
    ADD COLUMN sales_rep CHAR(4) REFERENCES sales_reps(rep_code);

-- Update orders.sales_rep to reference sales_reps(rep_code)
ALTER TABLE orders
    DROP COLUMN sales_rep,
    ADD COLUMN sales_rep CHAR(4) REFERENCES sales_reps(rep_code);

-- +goose Down
-- Drop sales_rep column from customer_locations table
ALTER TABLE customer_locations
    DROP COLUMN sales_rep;

-- Drop sales_rep column from orders table
ALTER TABLE orders
    DROP COLUMN sales_rep;

-- Add sales_rep column back to customers table
ALTER TABLE customers
    ADD COLUMN sales_rep TEXT;