
-- name: GetUser :one
SELECT * FROM users WHERE username = $1;


-- name: CreateUser :one
WITH userid AS (
  INSERT INTO users (username, passwordHash) VALUES ($1, $2) RETURNING id
)
SELECT * FROM users WHERE id = (SELECT * FROM userid);