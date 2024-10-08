// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createDocument = `-- name: CreateDocument :one

INSERT INTO documents (user_id, document_type, document_number, first_name, last_name, sex, birth_date, expire_date, country_code)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, user_id, document_type, document_number, first_name, last_name, sex, birth_date, expire_date, country_code, created_at, updated_at
`

type CreateDocumentParams struct {
	UserID         pgtype.Int4
	DocumentType   string
	DocumentNumber string
	FirstName      pgtype.Text
	LastName       pgtype.Text
	Sex            pgtype.Text
	BirthDate      pgtype.Date
	ExpireDate     pgtype.Date
	CountryCode    pgtype.Text
}

// Documents
func (q *Queries) CreateDocument(ctx context.Context, arg CreateDocumentParams) (Document, error) {
	row := q.db.QueryRow(ctx, createDocument,
		arg.UserID,
		arg.DocumentType,
		arg.DocumentNumber,
		arg.FirstName,
		arg.LastName,
		arg.Sex,
		arg.BirthDate,
		arg.ExpireDate,
		arg.CountryCode,
	)
	var i Document
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DocumentType,
		&i.DocumentNumber,
		&i.FirstName,
		&i.LastName,
		&i.Sex,
		&i.BirthDate,
		&i.ExpireDate,
		&i.CountryCode,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createEmailVerificationToken = `-- name: CreateEmailVerificationToken :one

INSERT INTO email_verification_tokens (user_id, token, expires_at)
VALUES ($1, $2, $3)
RETURNING id, user_id, token, expires_at, created_at
`

type CreateEmailVerificationTokenParams struct {
	UserID    pgtype.Int4
	Token     string
	ExpiresAt pgtype.Timestamp
}

// Email Verification Tokens
func (q *Queries) CreateEmailVerificationToken(ctx context.Context, arg CreateEmailVerificationTokenParams) (EmailVerificationToken, error) {
	row := q.db.QueryRow(ctx, createEmailVerificationToken, arg.UserID, arg.Token, arg.ExpiresAt)
	var i EmailVerificationToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const createPayment = `-- name: CreatePayment :one

INSERT INTO payments (user_id, amount, credits_added, payment_status)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, amount, credits_added, payment_date, payment_status
`

type CreatePaymentParams struct {
	UserID        pgtype.Int4
	Amount        pgtype.Numeric
	CreditsAdded  int32
	PaymentStatus string
}

// Payments
func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRow(ctx, createPayment,
		arg.UserID,
		arg.Amount,
		arg.CreditsAdded,
		arg.PaymentStatus,
	)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.CreditsAdded,
		&i.PaymentDate,
		&i.PaymentStatus,
	)
	return i, err
}

const createScanningHistory = `-- name: CreateScanningHistory :one

INSERT INTO scanning_history (user_id, document_id, document_url)
VALUES ($1, $2, $3)
RETURNING id, user_id, document_id, document_url, scan_date
`

type CreateScanningHistoryParams struct {
	UserID      pgtype.Int4
	DocumentID  pgtype.Int4
	DocumentUrl string
}

// Scanning History
func (q *Queries) CreateScanningHistory(ctx context.Context, arg CreateScanningHistoryParams) (ScanningHistory, error) {
	row := q.db.QueryRow(ctx, createScanningHistory, arg.UserID, arg.DocumentID, arg.DocumentUrl)
	var i ScanningHistory
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DocumentID,
		&i.DocumentUrl,
		&i.ScanDate,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one

INSERT INTO users (email, password_hash)
VALUES ($1, $2)
RETURNING id, email, password_hash, is_email_verified, credits, created_at, updated_at
`

type CreateUserParams struct {
	Email        string
	PasswordHash string
}

// Users
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Email, arg.PasswordHash)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.IsEmailVerified,
		&i.Credits,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUserSession = `-- name: CreateUserSession :one

INSERT INTO user_sessions (user_id, token, expires_at)
VALUES ($1, $2, $3)
RETURNING id, user_id, token, expires_at, created_at
`

type CreateUserSessionParams struct {
	UserID    pgtype.Int4
	Token     string
	ExpiresAt pgtype.Timestamp
}

// User Sessions (Optional)
func (q *Queries) CreateUserSession(ctx context.Context, arg CreateUserSessionParams) (UserSession, error) {
	row := q.db.QueryRow(ctx, createUserSession, arg.UserID, arg.Token, arg.ExpiresAt)
	var i UserSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const getDocumentsByUserId = `-- name: GetDocumentsByUserId :many
SELECT id, user_id, document_type, document_number, first_name, last_name, sex, birth_date, expire_date, country_code, created_at, updated_at
FROM documents
WHERE user_id = $1
`

func (q *Queries) GetDocumentsByUserId(ctx context.Context, userID pgtype.Int4) ([]Document, error) {
	rows, err := q.db.Query(ctx, getDocumentsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Document
	for rows.Next() {
		var i Document
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.DocumentType,
			&i.DocumentNumber,
			&i.FirstName,
			&i.LastName,
			&i.Sex,
			&i.BirthDate,
			&i.ExpireDate,
			&i.CountryCode,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEmailVerificationToken = `-- name: GetEmailVerificationToken :one
SELECT id, user_id, token, expires_at, created_at
FROM email_verification_tokens
WHERE token = $1
`

func (q *Queries) GetEmailVerificationToken(ctx context.Context, token string) (EmailVerificationToken, error) {
	row := q.db.QueryRow(ctx, getEmailVerificationToken, token)
	var i EmailVerificationToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const getPaymentsByUserId = `-- name: GetPaymentsByUserId :many
SELECT id, user_id, amount, credits_added, payment_date, payment_status
FROM payments
WHERE user_id = $1
`

func (q *Queries) GetPaymentsByUserId(ctx context.Context, userID pgtype.Int4) ([]Payment, error) {
	rows, err := q.db.Query(ctx, getPaymentsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			&i.CreditsAdded,
			&i.PaymentDate,
			&i.PaymentStatus,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getScanningHistoryByUserId = `-- name: GetScanningHistoryByUserId :many
SELECT id, user_id, document_id, document_url, scan_date
FROM scanning_history
WHERE user_id = $1
`

func (q *Queries) GetScanningHistoryByUserId(ctx context.Context, userID pgtype.Int4) ([]ScanningHistory, error) {
	rows, err := q.db.Query(ctx, getScanningHistoryByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ScanningHistory
	for rows.Next() {
		var i ScanningHistory
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.DocumentID,
			&i.DocumentUrl,
			&i.ScanDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password_hash, is_email_verified, credits, created_at, updated_at
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.IsEmailVerified,
		&i.Credits,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserSession = `-- name: GetUserSession :one
SELECT id, user_id, token, expires_at, created_at
FROM user_sessions
WHERE token = $1
`

func (q *Queries) GetUserSession(ctx context.Context, token string) (UserSession, error) {
	row := q.db.QueryRow(ctx, getUserSession, token)
	var i UserSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserCredits = `-- name: UpdateUserCredits :exec
UPDATE users
SET credits = $1
WHERE id = $2
`

type UpdateUserCreditsParams struct {
	Credits pgtype.Int4
	ID      int32
}

func (q *Queries) UpdateUserCredits(ctx context.Context, arg UpdateUserCreditsParams) error {
	_, err := q.db.Exec(ctx, updateUserCredits, arg.Credits, arg.ID)
	return err
}

const verifyEmail = `-- name: VerifyEmail :exec
UPDATE users
SET is_email_verified = TRUE
WHERE id = $1
`

func (q *Queries) VerifyEmail(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, verifyEmail, id)
	return err
}
