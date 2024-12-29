-- name: CreateOTP :exec
INSERT INTO otps (email, code, created_at, expires_at)
VALUES (?, ?, (datetime(CURRENT_TIMESTAMP, 'localtime')), ?);

-- name: GetOTPByEmailAndCode :one
SELECT *
FROM otps
WHERE email = ? AND code = ?;

-- name: DeleteOTPByEmail :exec
DELETE FROM otps
WHERE email = ?;