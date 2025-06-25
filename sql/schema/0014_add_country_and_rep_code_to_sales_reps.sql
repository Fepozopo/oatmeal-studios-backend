-- +goose Up
ALTER TABLE sales_reps
    ADD COLUMN country CHAR(3) NOT NULL DEFAULT 'USA',
    ADD COLUMN rep_code CHAR(4) NOT NULL UNIQUE;

-- +goose Down
ALTER TABLE sales_reps
    DROP COLUMN country,
    DROP COLUMN rep_code;
