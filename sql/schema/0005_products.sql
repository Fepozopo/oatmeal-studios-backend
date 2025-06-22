-- +goose Up
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    type TEXT NOT NULL,
    sku TEXT UNIQUE NOT NULL,
    upc TEXT UNIQUE NOT NULL,
    status TEXT NOT NULL,
    cost FLOAT NOT NULL,
    price FLOAT NOT NULL,
    envelope TEXT,
    artist TEXT,
    category TEXT,
    release_date DATE,
    last_bought_date DATE,
    description TEXT,
    text_front TEXT,
    text_inside TEXT
);

-- +goose Down
DROP TABLE products;
