-- +goose Up
CREATE TABLE inventory_transactions (
    id SERIAL PRIMARY KEY,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    change INTEGER NOT NULL, -- positive for Restock, negative for Sale/Adjustment
    reason TEXT NOT NULL,    -- e.g., 'Sale', 'Restock', 'Adjustment'
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
);

-- +goose Down
DROP TABLE inventory_transactions;