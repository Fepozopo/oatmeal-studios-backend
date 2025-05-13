-- +goose Up
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    business_name TEXT NOT NULL,
    contact_name TEXT, -- Optional contact person within the business
    email TEXT,
    phone TEXT,
    address_1 TEXT,
    address_2 TEXT,
    city TEXT,
    state TEXT,
    zip_code TEXT,
    terms TEXT,
    discount DECIMAL NOT NULL DEFAULT 0.0, -- Percentage discount
    commission DECIMAL NOT NULL DEFAULT 0.0, -- Percentage commission
    notes TEXT,
    sales_rep TEXT
);

-- +goose Down
DROP TABLE customers;
