package models

import (
	"database/sql"
	"testing"

	"github.com/icrowley/fake"
	"github.com/jmataya/hermes/config"
	"github.com/jmataya/hermes/errors"
	_ "github.com/lib/pq"
)

func TestCreateUser(t *testing.T) {
	dbConfig, err := config.NewDB()
	if err != nil {
		t.Error(err)
		return
	}

	dsn := dbConfig.PostgresDSN()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Error(err)
		return
	}

	defer db.Close()

	commonDB, err := NewDB(db)
	if err != nil {
		t.Error(err)
		return
	}

	var tests = []struct {
		model Model
		want  error
	}{
		{
			model: &User{
				id:        0,
				Email:     fake.EmailAddress(),
				FirstName: fake.FirstName(),
				LastName:  fake.LastName(),
			},
			want: nil,
		},
		{
			model: &User{
				id:        0,
				Email:     "",
				FirstName: fake.FirstName(),
				LastName:  fake.LastName(),
			},
			want: errors.NewFieldIsNil("email"),
		},
		{
			model: &User{
				id:        1,
				Email:     fake.EmailAddress(),
				FirstName: fake.FirstName(),
				LastName:  fake.LastName(),
			},
			want: errors.NewModelHasID("user", 1),
		},
	}

	for _, test := range tests {
		err := test.model.Create(commonDB)
		checkErr(t, "Failed to create user", test.want, err)
	}
}

func TestUpdateUser(t *testing.T) {
	dbConfig, err := config.NewDB()
	if err != nil {
		t.Error(err)
		return
	}

	dsn := dbConfig.PostgresDSN()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Error(err)
		return
	}

	defer db.Close()

	commonDB, err := NewDB(db)
	if err != nil {
		t.Error(err)
		return
	}

	user := &User{
		Email:     fake.EmailAddress(),
		FirstName: fake.FirstName(),
		LastName:  fake.LastName(),
	}

	err = user.Create(commonDB)
	checkErr(t, "Failed to create user", nil, err)

	user.FirstName = "New"
	err = user.Update(commonDB)
	checkErr(t, "Failed to update user", nil, err)
}

func checkErr(t *testing.T, msg string, want error, got error) {
	if want == got {
		return
	}

	if want == nil {
		t.Errorf("%s without error, got %s", msg, got.Error())
	} else if got == nil {
		t.Errorf("%s want %s, got no error", msg, want.Error())
	} else if want.Error() != got.Error() {
		t.Errorf("%s want %s, got %s", msg, want.Error(), got.Error())
	}
}
