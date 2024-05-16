
-- name: GetUser :one
SELECT * FROM users WHERE username = $1;


-- name: CreateUser :one
INSERT INTO users (username, passwordHash) VALUES ($1, $2) RETURNING *;