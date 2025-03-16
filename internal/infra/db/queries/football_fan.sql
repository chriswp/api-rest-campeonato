-- name: FindFootballFanByEmail :one
SELECT * FROM football_fans WHERE email = $1;

-- name: CreateFootballFan :one
INSERT INTO football_fans (id, name, email, team) VALUES ($1, $2, $3, $4) RETURNING *;