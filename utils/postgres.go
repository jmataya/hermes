package utils

import (
	"fmt"

	"github.com/jmataya/hermes/errors"
	"github.com/mattes/migrate"

	// Imports needed for the migration tool to work.
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

// MigratePG updates an existing Postgres database to have the most
// up-to-date schema.
func MigratePG(sourceDir, host, user, password, database string, isSSL bool) error {
	// Validate inputs.
	if sourceDir == "" {
		return errors.NewFieldIsNil("sourceDir")
	} else if host == "" {
		return errors.NewFieldIsNil("host")
	} else if user == "" {
		return errors.NewFieldIsNil("user")
	} else if database == "" {
		return errors.NewFieldIsNil("database")
	}

	source := fmt.Sprintf("file://%s", sourceDir)

	pwd := ""
	if password != "" {
		pwd = fmt.Sprintf(":%s", password)
	}

	var ssl string
	if isSSL {
		ssl = "enable"
	} else {
		ssl = "disable"
	}

	conn := fmt.Sprintf("postgres://%s%s@%s:5432/%s?sslmode=%s", user, pwd, host, database, ssl)

	m, err := migrate.New(source, conn)
	if err != nil {
		return fmt.Errorf("Failed to initialize migrations withe error: %s", err.Error())
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return fmt.Errorf("Failed to run migrations with error: %s", err.Error())
	}

	return nil
}
