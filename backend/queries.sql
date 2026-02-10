-- name: GetAllAthletes :many
SELECT id, name, grade, personal_record, events FROM athletes ORDER BY name;

-- name: GetAthleteByID :one
SELECT id, name, grade, personal_record, events FROM athletes WHERE id = ?;

-- name: GetAllMeets :many
SELECT id, name, date, location, description FROM meets ORDER BY date;

-- name: GetResultsByMeet :many
SELECT r.id, r.athlete_id, r.meet_id, r.time, r.place, a.name AS athlete_name
FROM results r
JOIN athletes a ON r.athlete_id = a.id
WHERE r.meet_id = ?
ORDER BY r.place;

-- name: CreateResult :one
INSERT INTO results (athlete_id, meet_id, time, place)
VALUES (?, ?, ?, ?)
RETURNING id, athlete_id, meet_id, time, place;

-- name: GetTopTimes :many
SELECT r.id, r.time, r.place, a.name AS athlete_name, m.name AS meet_name
FROM results r
JOIN athletes a ON r.athlete_id = a.id
JOIN meets m ON r.meet_id = m.id
ORDER BY r.time ASC
LIMIT 10;
