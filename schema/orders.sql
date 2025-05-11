CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    customer_id INTEGER NOT NULL REFERENCES customers(id),
    order_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    status TEXT NOT NULL,
    type TEXT NOT NULL, -- e.g., 'reorder', 'credit'
    method TEXT, -- e.g., 'online', 'phone', 'email'
    ship_date TIMESTAMPTZ,
    po_number TEXT,
    shipping_cost DECIMAL NOT NULL DEFAULT 0.0,
    free_shipping BOOLEAN NOT NULL DEFAULT FALSE,
    apply_to_commission BOOLEAN NOT NULL DEFAULT FALSE,
    notes TEXT
);
