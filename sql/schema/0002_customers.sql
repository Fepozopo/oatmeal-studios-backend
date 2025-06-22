-- +goose Up
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    business_name TEXT NOT NULL,
    contact_name TEXT, -- Optional contact person within the business
    email TEXT,
    phone TEXT,
    address_1 TEXT NOT NULL,
    address_2 TEXT,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    zip_code TEXT NOT NULL,
    terms TEXT NOT NULL DEFAULT 'Net 30', -- Payment terms
    discount FLOAT NOT NULL DEFAULT 0.0, -- Percentage discount
    commission FLOAT NOT NULL DEFAULT 0.0, -- Percentage commission
    sales_rep TEXT,
    notes TEXT
);

-- +goose Down
DROP TABLE customers;
