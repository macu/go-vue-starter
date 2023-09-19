package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

// DBConn wraps a *sql.DB or *sql.Tx.
type DBConn interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// Returns a placeholder representing the given arg in args,
// adding the arg to args if not already present.
func argPlaceholder(arg interface{}, args *[]interface{}) string {
	for i := 0; i < len(*args); i++ {
		if (*args)[i] == arg {
			return "$" + strconv.Itoa(i+1)
		}
	}
	*args = append(*args, arg)
	return "$" + strconv.Itoa(len(*args))
}

// Returns the "= $N" or "IS NULL" part of an equality condition
// where the operand may be null.
func eqCond(col string, arg interface{}, args *[]interface{}) string {
	if isNil(arg) {
		return col + " IS NULL"
	}
	return col + " = " + argPlaceholder(arg, args)
}

// Returns the "= $N" or "IS NULL" part of an equality condition
// where the operand may be null, and the placeholder index is given.
func eqCondIndexed(col string, arg interface{}, index int) string {
	if isNil(arg) {
		return col + " IS NULL"
	}
	return col + " = $" + strconv.Itoa(index)
}

func createArgsList(args *[]interface{}, values ...interface{}) string {
	out := ``
	for i := 0; i < len(values); i++ {
		if i > 0 {
			out += `,`
		}
		out += argPlaceholder(values[i], args)
	}
	return out
}

func createArgsListInt64s(args *[]interface{}, values ...int64) string {
	out := ``
	for i := 0; i < len(values); i++ {
		if i > 0 {
			out += `,`
		}
		out += argPlaceholder(values[i], args)
	}
	return out
}

// create a VALUES (), (), ... postgres string using argument placeholders
func argValuesMap(args *[]interface{}, values [][]interface{}) string {
	var out = `VALUES `

	for i := 0; i < len(values); i++ {
		if i > 0 {
			out += `,`
		}

		out += `(`

		for j := 0; j < len(values[i]); j++ {
			if j > 0 {
				out += `,`
			}

			var v = values[i][j]
			out += argPlaceholder(v, args)

			// include type casts with placeholders
			switch v.(type) {
			case int, uint, int64:
				out += `::int`
			case string:
				out += `::text`
			}
		}

		out += `)`
	}

	return out
}

func inTransaction(r *http.Request, db *sql.DB, f func(*sql.Tx) error) error {
	// Thanks GPT-3.5 for guidance

	c := r.Context()
	tx, err := db.BeginTx(c, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	err = f(tx)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("rollback: %v; on run function: %w", rbErr, err)
		}
		return fmt.Errorf("run function: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit: %w", err)
	}

	return nil
}

func handleInTransaction(r *http.Request, db *sql.DB, userID *uint,
	f func(*sql.Tx) (interface{}, int, error)) (interface{}, int) {
	// Thanks GPT-3.5 for guidance

	c := r.Context()
	tx, err := db.BeginTx(c, nil)
	if err != nil {
		logError(r, userID, fmt.Errorf("begin transaction: %w", err))
		return nil, http.StatusInternalServerError
	}

	response, statusCode, err := f(tx)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			logError(r, userID, fmt.Errorf("rollback: %v; on run function: %w", rbErr, err))
			return nil, http.StatusInternalServerError
		}
		logError(r, userID, fmt.Errorf("run function: %w", err))
		return nil, statusCode
	}

	err = tx.Commit()
	if err != nil {
		logError(r, userID, fmt.Errorf("commit: %w", err))
		return nil, http.StatusInternalServerError
	}

	return response, statusCode
}
