-- +goose Up
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    sku TEXT UNIQUE NOT NULL REFERENCES products(sku) ON DELETE CASCADE,
    quantity INTEGER NOT NULL DEFAULT 6,
    price FLOAT NOT NULL,
    discount FLOAT NOT NULL DEFAULT 0.0, -- Percentage discount on the item
    item_total FLOAT NOT NULL,
    pocket_number INTEGER -- If item is on the customer location's planogram, this is the pocket number
);

-- +goose Down
DROP TABLE order_items;
