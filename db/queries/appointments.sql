-- name: CreateAppointment :exec
INSERT INTO appointments (date_time, car_name, status, user_id, service_type) VALUES ($1, $2, $3, $4, $5);

-- name: GetAllAppointments :many
SELECT * FROM appointments;

-- name: GetAppointmentByID :one
SELECT * FROM appointments where id = $1;

-- name: GetAppointmentsByUserID :many
SELECT * FROM appointments where user_id = $1;

