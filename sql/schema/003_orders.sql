
-- +goose Up
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    customer_id INTEGER NOT NULL REFERENCES customers(id),
    customer_location_id INTEGER REFERENCES customer_locations(id),
    order_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ship_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status TEXT NOT NULL,
    type TEXT NOT NULL, -- e.g., 'reorder', 'credit'
    method TEXT, -- e.g., 'online', 'phone', 'email'
    po_number TEXT,
    shipping_cost FLOAT NOT NULL DEFAULT 0.0,
    free_shipping BOOLEAN NOT NULL DEFAULT FALSE,
    apply_to_commission BOOLEAN NOT NULL DEFAULT FALSE,
    sales_rep TEXT,
    notes TEXT
);

-- +goose Down
DROP TABLE orders;