package models

import (
	"database/sql"

	"github.com/jmataya/hermes/errors"
)

// DB is a wrapper around the standard database/sql interface that provides
// error handling logic, allowing for easier code maintenance in application
// code.
type DB interface {
	Error() error
	Prepare(string) *sql.Stmt
	QueryRow(string, ...interface{}) *sql.Row

	RowScan(*sql.Row, ...interface{})

	StmtExec(*sql.Stmt, ...interface{}) sql.Result
}

// NewDB creates an new DB from an existing sql/db.
func NewDB(db *sql.DB) (DB, error) {
	if db == nil {
		return nil, errors.NewFieldIsNil("db")
	}

	return &sqlDB{db: db}, nil
}

type sqlDB struct {
	db  *sql.DB
	err error
}

func (s *sqlDB) Error() error {
	return s.err
}

func (s *sqlDB) Prepare(query string) *sql.Stmt {
	if s.err != nil {
		return nil
	}

	stmt, err := s.db.Prepare(query)
	if err != nil {
		s.err = err
		return nil
	}

	return stmt
}

func (s *sqlDB) QueryRow(query string, args ...interface{}) *sql.Row {
	if s.err != nil {
		return nil
	}

	return s.db.QueryRow(query, args...)
}

func (s *sqlDB) RowScan(row *sql.Row, dest ...interface{}) {
	if s.err != nil {
		return
	}

	err := row.Scan(dest...)
	s.err = err
}

func (s *sqlDB) StmtExec(stmt *sql.Stmt, args ...interface{}) sql.Result {
	if s.err != nil {
		return nil
	}

	res, err := stmt.Exec(args...)
	if err != nil {
		s.err = err
		return nil
	}

	return res
}
