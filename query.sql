
-- name: GetUser :one
SELECT * FROM users WHERE username = $1;


-- name: CreateUser :one
WITH userid AS (
  INSERT INTO users (username, passwordHash) VALUES ($1, $2) RETURNING id
)
SELECT * FROM users WHERE id = (SELECT * FROM userid);

-- name: CreateClub :exec
WITH clubID AS (
    INSERT INTO clubs (name, description) VALUES ($1, $2) RETURNING id
)
INSERT INTO managers (user_id, club_id) VALUES ($3, clubID);


-- name: addManager :exec
INSERT INTO managers (user_id, club_id) VALUES ($1, $2);

-- name: CreateEvent :exec
WITH eventID AS (
    INSERT INTO events (name, place, description, start_time, end_time, club_id) 
        VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
)
INSERT INTO organizers (user_id, event_id) VALUES ($7, eventID);

-- name: addOrganizer :exec
INSERT INTO organizers (user_id, event_id) VALUES ($1, $2);

-- name: GetClubs :many
SELECT * FROM clubs;

-- name: GetClub :one
SELECT * FROM clubs WHERE id = $1;

-- name: GetFutureEvents :many
SELECT * FROM events WHERE start_time > NOW();

-- name: GetEvent :one
SELECT * FROM events WHERE id = $1;

-- name: GetOrganizers :many
SELECT * FROM organizers WHERE event_id = $1;

-- name: GetManagers :many
SELECT * FROM managers WHERE club_id = $1;


