-- +goose Up
ALTER TABLE planogram_pockets
    DROP CONSTRAINT IF EXISTS planogram_pockets_product_id_fkey;

ALTER TABLE planogram_pockets
    DROP COLUMN IF EXISTS product_id;

ALTER TABLE planogram_pockets
    ADD COLUMN sku TEXT REFERENCES products(sku);

-- +goose Down
ALTER TABLE planogram_pockets
    DROP COLUMN IF EXISTS sku;

ALTER TABLE planogram_pockets
    ADD COLUMN product_id UUID REFERENCES products(id);
