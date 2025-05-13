
-- +goose Up
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    customer_id INTEGER NOT NULL REFERENCES customers(id),
    order_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ship_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status TEXT NOT NULL,
    type TEXT NOT NULL, -- e.g., 'reorder', 'credit'
    method TEXT, -- e.g., 'online', 'phone', 'email'
    po_number TEXT,
    shipping_cost DECIMAL NOT NULL DEFAULT 0.0,
    free_shipping BOOLEAN NOT NULL DEFAULT FALSE,
    apply_to_commission BOOLEAN NOT NULL DEFAULT FALSE,
    notes TEXT,
    sales_rep TEXT
);

-- +goose Down
DROP TABLE orders;