
-- name: GetCurrentInventory :one
SELECT COALESCE(SUM(change), 0) AS current_inventory
FROM inventory_transactions
WHERE product_id = $1;

-- name: GetAllCurrentInventory :many
SELECT
  product_id,
  COALESCE(SUM(change), 0) AS current_inventory
FROM inventory_transactions
GROUP BY product_id;

-- name: ListInventoryTransactions :many
SELECT *
FROM inventory_transactions
WHERE product_id = $1
ORDER BY created_at DESC;

-- name: InsertInventoryTransaction :one
INSERT INTO inventory_transactions (product_id, change, reason, notes)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetInventoryChangesByDay :many
SELECT
  DATE(created_at) AS day,
  SUM(change) AS inventory_change
FROM inventory_transactions
WHERE product_id = $1
GROUP BY day
ORDER BY day;
