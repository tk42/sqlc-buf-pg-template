// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: pet.sql

package petapis

import (
	"context"
	"database/sql"
)

const createPetQuery = `-- name: CreatePetQuery :one
INSERT INTO pets (
  name, memo
) VALUES (
  $1, $2
)
RETURNING id, name, memo, create_at
`

type CreatePetQueryParams struct {
	Name string
	Memo sql.NullString
}

func (q *Queries) CreatePetQuery(ctx context.Context, arg CreatePetQueryParams) (Pet, error) {
	row := q.db.QueryRowContext(ctx, createPetQuery, arg.Name, arg.Memo)
	var i Pet
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Memo,
		&i.CreateAt,
	)
	return i, err
}

const deletePetQuery = `-- name: DeletePetQuery :exec
DELETE FROM pets
WHERE id = $1
`

func (q *Queries) DeletePetQuery(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePetQuery, id)
	return err
}

const getPetQuery = `-- name: GetPetQuery :one
SELECT id, name, memo, create_at FROM pets
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPetQuery(ctx context.Context, id int64) (Pet, error) {
	row := q.db.QueryRowContext(ctx, getPetQuery, id)
	var i Pet
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Memo,
		&i.CreateAt,
	)
	return i, err
}

const listPetsQuery = `-- name: ListPetsQuery :many
SELECT id, name, memo, create_at FROM pets
ORDER BY name
`

func (q *Queries) ListPetsQuery(ctx context.Context) ([]Pet, error) {
	rows, err := q.db.QueryContext(ctx, listPetsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pet
	for rows.Next() {
		var i Pet
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Memo,
			&i.CreateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}