package models

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/jmataya/hermes/errors"
	_ "github.com/lib/pq"
)

func TestCreateUser(t *testing.T) {
	db, err := InitializeDB()
	if err != nil {
		t.Error(err)
		return
	}

	defer db.Close()

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
		err := test.model.Create(db)
		checkErr(t, "Failed to create user", test.want, err)
	}
}

func TestUpdateUser(t *testing.T) {
	db, err := InitializeDB()
	if err != nil {
		t.Error(err)
		return
	}

	defer db.Close()

	user := &User{
		Email:     fake.EmailAddress(),
		FirstName: fake.FirstName(),
		LastName:  fake.LastName(),
	}

	err = user.Create(db)
	checkErr(t, "Failed to create user", nil, err)

	user.FirstName = "New"
	err = user.Update(db)
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
