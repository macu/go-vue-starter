package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Username pattern
var usernameRegex = regexp.MustCompile("^[a-zA-Z0-9_]+$")

// Regex adapted from https://www.w3.org/TR/html5/forms.html#valid-e-mail-address
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

const usernameMaxLength = 25
const emailAddressMaxLength = 50
const passwordMinLength = 5

// BcryptCost is the cost applied to hashing and subsequently verifying new passwords.
const bcryptCost = 15

// Returns the user ID if the user was successfully created.
func createUser(r *http.Request, db *sql.DB,
	username, password, email string) (uint, error) {

	var newUserID uint
	var err error

	err = inTransaction(r, db, func(tx *sql.Tx) error {
		// inTransaction may return same or different error
		newUserID, err = createUserTx(tx, username, password, email)
		return err
	})

	if err != nil {
		return 0, err
	}

	return newUserID, nil
}

// Creates a user within an existing transaction.
// Error is returned on failure; cancel the transaction in the parent context.
// Returns the user ID if the user was successfully created.
func createUserTx(tx *sql.Tx, username, password, email string) (uint, error) {

	username = strings.TrimSpace(username)
	email = strings.TrimSpace(email)

	if username == "" {
		return 0, fmt.Errorf("username required")
	}
	if len(username) > usernameMaxLength {
		return 0, fmt.Errorf("username must be %d characters or less", usernameMaxLength)
	}
	if email == "" {
		return 0, fmt.Errorf("email required")
	}
	if len(email) > emailAddressMaxLength {
		return 0, fmt.Errorf("email must be %d characters or less", emailAddressMaxLength)
	}
	if !emailRegexp.MatchString(email) {
		return 0, fmt.Errorf("invalid email address")
	}
	if len(strings.TrimSpace(password)) < passwordMinLength {
		return 0, fmt.Errorf("password must be %d characters or more", passwordMinLength)
	}

	var exists bool
	err := tx.QueryRow(
		`SELECT EXISTS(SELECT * FROM user_account WHERE username = $1)`, username,
	).Scan(&exists)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, fmt.Errorf("username already exists")
	}

	authHash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return 0, err
	}

	var userID uint
	err = tx.QueryRow(
		`INSERT INTO user_account (username, email, auth_hash, created_at)
		VALUES ($1, $2, $3, $4) RETURNING id`,
		username, email, authHash, time.Now()).Scan(&userID)
	if err != nil {
		return 0, err
	}

	// TODO Create any initial user-related records in same transaction

	return userID, nil
}
