-- name: GetPetQuery :one
SELECT * FROM pets
WHERE id = $1 LIMIT 1;

-- name: ListPetsQuery :many
SELECT * FROM pets
ORDER BY name;

-- name: CreatePetQuery :one
INSERT INTO pets (
  name, memo
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeletePetQuery :exec
DELETE FROM pets
WHERE id = $1;