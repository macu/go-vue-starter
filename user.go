package main

import (
	"database/sql"
	"errors"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Regex adapted from https://www.w3.org/TR/html5/forms.html#valid-e-mail-address
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Returns the user ID if the user was successfully created.
func createUser(db *sql.DB, username, password, email string) (int64, error) {
	username = strings.TrimSpace(username)
	email = strings.TrimSpace(email)
	if username == "" {
		return 0, errors.New("Username required")
	}
	if len(username) > 25 {
		return 0, errors.New("Username must be 25 characters or less")
	}
	if email == "" {
		return 0, errors.New("Email required")
	}
	if len(email) > 50 {
		return 0, errors.New("Email must be 50 characters or less")
	}
	if !emailRegexp.MatchString(email) {
		return 0, errors.New("Invalid email address")
	}
	if password == "" {
		return 0, errors.New("Password must not be empty")
	}

	existing := db.QueryRow("SELECT EXISTS(SELECT * FROM user WHERE username=?)", username)
	var exists bool
	err := existing.Scan(&exists)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, errors.New("Username already exists")
	}

	authHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	res, err := tx.Exec("INSERT INTO user (username, email, auth_hash, created_at) VALUES (?, ?, ?, ?)",
		username, email, authHash, time.Now())
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	userID, err := res.LastInsertId()
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	// TODO Create any initial user-related records in same transaction

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return userID, nil
}
