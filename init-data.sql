


-- username: test1 password: Test1234%
INSERT INTO users (id, username, passwordHash)
VALUES (1, 'test1', '$2a$10$aGqwT7T6aOGcuDnLaavKvOtBlWe2XwCU8ykQmNqEo7bFVshjYFcBC');

-- username: test2 password: Test1234%
INSERT INTO users (id, username, passwordHash)
VALUES (2, 'test2', '$2a$10$aGqwT7T6aOGcuDnLaavKvOtBlWe2XwCU8ykQmNqEo7bFVshjYFcBC');


-- club: club1
WITH clubID AS (
    INSERT INTO clubs (name, description) VALUES ('club1', 'a club') RETURNING id
)
INSERT INTO managers (user_id, club_id)
SELECT 1, id FROM clubID;

-- club: club2
WITH clubID AS (
    INSERT INTO clubs (name, description) VALUES ('club2', 'a second club') RETURNING id
)
INSERT INTO managers (user_id, club_id)
SELECT 1, id FROM clubID UNION
SELECT 2, id FROM clubID;


-- event: event1
WITH eventID AS (
    INSERT INTO events (name, place, description, start_time, end_time, club_id) 
        VALUES ('event1', 'place1', 'an event', '2030-01-01 00:00:00', '2030-01-01 01:00:00', 1) 
        RETURNING id
)
INSERT INTO organizers (user_id, event_id)  
SELECT 1, id FROM eventID;

-- event: event2
WITH eventID AS (
    INSERT INTO events (name, place, description, start_time, end_time, club_id) 
        VALUES ('event2', 'place2', 'another event', '2030-01-01 00:00:00', '2030-01-01 01:00:00', 2) 
        RETURNING id
)
INSERT INTO organizers (user_id, event_id)
SELECT 1, id FROM eventID;

-- event: event3
WITH eventID AS (
    INSERT INTO events (name, place, description, start_time, end_time, club_id) 
        VALUES ('event3', 'place3', 'yet another event', '2030-01-01 00:00:00', '2030-01-01 01:00:00', NULL) 
        RETURNING id
)
INSERT INTO organizers (user_id, event_id)
SELECT 1, id FROM eventID;


-- event: event4
WITH eventID AS (
    INSERT INTO events (name, place, description, start_time, end_time, club_id) 
        VALUES ('event4', 'place4', 'yet another event', '2030-01-01 00:00:00', '2030-01-01 01:00:00', 2) 
        RETURNING id
)
INSERT INTO organizers (user_id, event_id)
SELECT 1, id FROM eventID UNION
SELECT 2, id FROM eventID;


-- event: event5
WITH eventID AS (
    INSERT INTO events (name, place, description, start_time, end_time, club_id) 
        VALUES ('event5', 'place5', 'yet another event', '2030-01-01 00:00:00', '2030-01-01 01:00:00', NULL) 
        RETURNING id
)
INSERT INTO organizers (user_id, event_id)
SELECT 2, id FROM eventID;

