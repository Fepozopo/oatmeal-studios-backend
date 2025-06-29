-- +goose Up
ALTER TABLE sales_reps ALTER COLUMN country DROP NOT NULL;

-- +goose Down
ALTER TABLE sales_reps ALTER COLUMN country SET NOT NULL;
