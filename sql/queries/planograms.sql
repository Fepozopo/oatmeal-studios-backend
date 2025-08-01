-- name: GetPlanogram :one
SELECT * FROM planograms WHERE id = $1;

-- name: ListPlanograms :many
SELECT * FROM planograms ORDER BY created_at DESC;

-- name: CreatePlanogram :one
INSERT INTO planograms (name, num_pockets, notes)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdatePlanogram :one
UPDATE planograms
SET name = $2,
    num_pockets = $3,
    notes = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePlanogram :exec
DELETE FROM planograms WHERE id = $1;

-- name: AssignPlanogramToLocation :one
INSERT INTO planogram_customer_locations (planogram_id, customer_location_id)
VALUES ($1, $2)
ON CONFLICT (planogram_id, customer_location_id) DO NOTHING
RETURNING *;

-- name: ReassignPlanogramToLocation :one
WITH deleted AS (
  DELETE FROM planogram_customer_locations
  WHERE customer_location_id = $2
)
INSERT INTO planogram_customer_locations (planogram_id, customer_location_id)
VALUES ($1, $2)
RETURNING id, planogram_id, customer_location_id;

-- name: RemovePlanogramFromLocation :exec
DELETE FROM planogram_customer_locations
WHERE customer_location_id = $1;

-- name: GetPlanogramsByLocation :one
SELECT p.*
FROM planograms p
JOIN planogram_customer_locations pcl ON p.id = pcl.planogram_id
WHERE pcl.customer_location_id = $1;

-- name: ListLocationsByPlanogram :many
SELECT cl.*
FROM customer_locations cl
JOIN planogram_customer_locations pcl ON cl.id = pcl.customer_location_id
WHERE pcl.planogram_id = $1
ORDER BY cl.id;

-- name: ListPocketsForPlanogram :many
SELECT * FROM planogram_pockets WHERE planogram_id = $1 ORDER BY pocket_number;

-- name: GetPlanogramPocket :one
SELECT * FROM planogram_pockets WHERE id = $1;

-- name: CreatePlanogramPocket :one
INSERT INTO planogram_pockets (planogram_id, pocket_number, category, sku)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdatePlanogramPocket :one
UPDATE planogram_pockets
SET category = $2,
    sku = $3
WHERE id = $1
RETURNING *;

-- name: DeletePlanogramPocket :exec
DELETE FROM planogram_pockets WHERE id = $1;

-- name: GetPlanogramPocketByNumber :one
SELECT * FROM planogram_pockets WHERE planogram_id = $1 AND pocket_number = $2;
