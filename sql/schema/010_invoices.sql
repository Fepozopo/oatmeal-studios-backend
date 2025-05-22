-- +goose Up
CREATE TABLE invoices (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    invoice_date TIMESTAMP NOT NULL DEFAULT NOW(),
    order_id INT NOT NULL,
    customer_id INT NOT NULL,
    customer_location_id INT,
    due_date TIMESTAMP NOT NULL,
    status VARCHAR(50) NOT NULL,
    total FLOAT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE invoices;