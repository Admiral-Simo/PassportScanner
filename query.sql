-- Users

-- name: CreateUser :one
INSERT INTO users (email, password_hash)
VALUES ($1, $2)
RETURNING id, email, password_hash, is_email_verified, credits, created_at, updated_at;

-- name: GetUserByEmail :one
SELECT id, email, password_hash, is_email_verified, credits, created_at, updated_at
FROM users
WHERE email = $1;

-- name: UpdateUserCredits :exec
UPDATE users
SET credits = $1
WHERE id = $2;

-- name: VerifyEmail :exec
UPDATE users
SET is_email_verified = TRUE
WHERE id = $1;

-- Documents

-- name: CreateDocument :one
INSERT INTO documents (user_id, document_type, document_number, first_name, last_name, sex, birth_date, expire_date, country_code)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, user_id, document_type, document_number, first_name, last_name, sex, birth_date, expire_date, country_code, created_at, updated_at;

-- name: GetDocumentsByUserId :many
SELECT id, user_id, document_type, document_number, first_name, last_name, sex, birth_date, expire_date, country_code, created_at, updated_at
FROM documents
WHERE user_id = $1;

-- Payments

-- name: CreatePayment :one
INSERT INTO payments (user_id, amount, credits_added, payment_status)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, amount, credits_added, payment_date, payment_status;

-- name: GetPaymentsByUserId :many
SELECT id, user_id, amount, credits_added, payment_date, payment_status
FROM payments
WHERE user_id = $1;

-- Email Verification Tokens

-- name: CreateEmailVerificationToken :one
INSERT INTO email_verification_tokens (user_id, token, expires_at)
VALUES ($1, $2, $3)
RETURNING id, user_id, token, expires_at, created_at;

-- name: GetEmailVerificationToken :one
SELECT id, user_id, token, expires_at, created_at
FROM email_verification_tokens
WHERE token = $1;

-- Scanning History

-- name: CreateScanningHistory :one
INSERT INTO scanning_history (user_id, document_id, document_url)
VALUES ($1, $2, $3)
RETURNING id, user_id, document_id, document_url, scan_date;

-- name: GetScanningHistoryByUserId :many
SELECT id, user_id, document_id, document_url, scan_date
FROM scanning_history
WHERE user_id = $1;

-- User Sessions (Optional)

-- name: CreateUserSession :one
INSERT INTO user_sessions (user_id, token, expires_at)
VALUES ($1, $2, $3)
RETURNING id, user_id, token, expires_at, created_at;

-- name: GetUserSession :one
SELECT id, user_id, token, expires_at, created_at
FROM user_sessions
WHERE token = $1;
