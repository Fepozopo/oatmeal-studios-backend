-- +goose Up
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    item TEXT NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 1,
    price DECIMAL NOT NULL,
    discount DECIMAL NOT NULL DEFAULT 0.0, -- Percentage discount on the item
    item_total DECIMAL NOT NULL,
    pocket_number INTEGER -- If item is on the customer location's planogram, this is the pocket number
);

-- +goose Down
DROP TABLE order_items;
