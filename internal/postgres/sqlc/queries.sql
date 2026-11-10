-- name: FindByCode :one
SELECT * FROM links WHERE code = $1;

-- name: Create :exec
INSERT INTO links(id, code, url) VALUES($1, $2, $3);

-- name: List :many
SELECT * FROM links;

-- name: FindByID :one
SELECT * FROM links WHERE id = $1;

-- name: Delete :exec
DELETE FROM links WHERE id = $1;