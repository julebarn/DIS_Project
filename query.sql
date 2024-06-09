
-- name: GetUser :one
SELECT * FROM users WHERE username = $1;


-- name: GetAllUsers :many
SELECT * FROM users;


-- name: CreateUser :one
INSERT INTO users (username, passwordHash) VALUES ($1, $2) RETURNING id;


-- name: CreateClub :exec
WITH clubID AS (
    INSERT INTO clubs (name, description) VALUES ($1, $2) RETURNING id
)
INSERT INTO managers (user_id, club_id)
SELECT $3, id FROM clubID;

-- name: AddManager :exec
INSERT INTO managers (user_id, club_id) VALUES ($1, $2);

-- name: CreateEvent :exec
WITH eventID AS (
    INSERT INTO events (name, place, description, start_time, end_time, club_id) 
        VALUES ($1, $2, $3, $4,$5, $6) 
        RETURNING id
)
INSERT INTO organizers (user_id, event_id) 
SELECT $7, id FROM eventID;

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


-- name: GetManagers :many
SELECT id, username FROM users WHERE id IN (SELECT user_id FROM managers WHERE club_id = $1);

-- name: GetOrganizers :many
SELECT id, username FROM users WHERE id IN (SELECT user_id FROM organizers WHERE event_id = $1);

-- name: GetClubByManagers :many
SELECT * FROM clubs WHERE id IN (SELECT club_id FROM managers WHERE user_id = $1);







