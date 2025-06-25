-- +goose Up
ALTER TABLE sales_reps ALTER COLUMN company DROP NOT NULL;

-- +goose Down
ALTER TABLE sales_reps ALTER COLUMN company SET NOT NULL;
