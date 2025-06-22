-- +goose Up
CREATE TABLE planograms (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    num_pockets INTEGER NOT NULL,
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE planogram_customer_locations (
    id SERIAL PRIMARY KEY,
    planogram_id INTEGER NOT NULL REFERENCES planograms(id) ON DELETE CASCADE,
    customer_location_id INTEGER NOT NULL REFERENCES customer_locations(id) ON DELETE CASCADE,
    UNIQUE(planogram_id, customer_location_id)
);

CREATE TABLE planogram_pockets (
    id SERIAL PRIMARY KEY,
    planogram_id INTEGER NOT NULL REFERENCES planograms(id) ON DELETE CASCADE,
    pocket_number INTEGER NOT NULL,
    category TEXT NOT NULL,
    product_id UUID REFERENCES products(id),
    UNIQUE(planogram_id, pocket_number)
);

-- +goose Down
DROP TABLE planogram_pockets;
DROP TABLE planogram_customer_locations;
DROP TABLE planograms;
