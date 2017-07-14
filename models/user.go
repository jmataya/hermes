package models

import (
	"time"

	"github.com/jmataya/hermes/errors"
)

// User represents an entity in the system that can interact with files, either
// by modifying or reading.
type User struct {
	id int

	Email     string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ID is the primary key of the model.
func (u *User) ID() int {
	return u.id
}

// Validate ensures that the model can be created or updated.
func (u *User) Validate() error {
	if u.Email == "" {
		return errors.NewFieldIsNil("email")
	}

	return nil
}

// Create inserts the user into the database.
func (u *User) Create(db DB) error {
	const insert = `
    INSERT INTO users (email, first_name, last_name, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *`

	if u.id != 0 {
		return errors.NewModelHasID("user", u.id)
	}

	if err := u.Validate(); err != nil {
		return err
	}

	row := db.QueryRow(
		insert,
		u.Email,
		u.FirstName,
		u.LastName,
		time.Now().UTC(),
		time.Now().UTC())

	db.RowScan(
		row,
		&u.id,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.CreatedAt,
		&u.UpdatedAt)

	return db.Error()
}

func (u *User) Update(db DB) error {
	const update = `
    UPDATE users
    SET email=$1, first_name=$2, last_name=$3, updated_at=$4
    WHERE id=$5
    RETURNING *`

	if u.id == 0 {
		return errors.NewRequiredFieldIsEmpty("user", "id")
	}

	row := db.QueryRow(
		update,
		u.Email,
		u.FirstName,
		u.LastName,
		time.Now().UTC(),
		u.id)

	db.RowScan(
		row,
		&u.id,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.CreatedAt,
		&u.UpdatedAt)

	return db.Error()
}

func (u *User) Delete(db DB) error {
	const delete = "DELETE FROM users WHERE id=$1"

	if u.id == 0 {
		return errors.NewRequiredFieldIsEmpty("user", "id")
	}

	if err := u.Validate(); err != nil {
		return err
	}

	stmt := db.Prepare(delete)
	res := db.StmtExec(stmt, u.id)

	if err := db.Error(); err != nil {
		return err
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowCount != 1 {
		return errors.NewDeleteFailed("user", u.id, rowCount)
	}

	return nil
}
