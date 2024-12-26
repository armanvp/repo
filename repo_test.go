package repo

import (
	"database/sql"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestRepo(t *testing.T) {
	r := require.New(t)

	database, err := sqlx.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
	r.NoError(err)

	tx, err := database.Beginx()
	r.NoError(err)

	query := `
		CREATE TABLE users (
			id INTEGER PRIMARY KEY,
			email text NOT NULL,
			first_name text NOT NULL,
			last_name text
		);
	`
	_, err = tx.Exec(query)
	r.NoError(err)
	err = tx.Commit()
	r.NoError(err)

	type User struct {
		Id        int            `db:"id"`
		Email     string         `db:"email"`
		FirstName string         `db:"first_name"`
		LastName  sql.NullString `db:"last_name"`
	}

	// Test Create
	tx, err = database.Beginx()
	r.NoError(err)

	repo := New[User]("users", tx)
	err = repo.Create(&User{
		Id:        40,
		Email:     "juan@delacruz.com",
		FirstName: "Juan",
		LastName: sql.NullString{
			String: "Delacruz",
			Valid:  true,
		},
	})
	r.NoError(err)

	err = tx.Commit()
	r.NoError(err)

	// Test Get
	tx, err = database.Beginx()
	r.NoError(err)

	repo = New[User]("users", tx)
	rec, err := repo.Get(40)
	r.NoError(err)
	r.Equal("juan@delacruz.com", rec.Email)
	r.Equal("Juan", rec.FirstName)
	r.Equal("Delacruz", rec.LastName.String)

	// Create new test data
	err = repo.Create(&User{
		Id:        41,
		Email:     "jane@delacruz.com",
		FirstName: "Jane",
		LastName: sql.NullString{
			String: "Delacruz",
			Valid:  true,
		},
	})
	r.NoError(err)

	// Test GetByParam
	recs, err := repo.GetByParam(Params{
		{Field: "last_name", Value: "Delacruz"},
	})
	r.NoError(err)
	r.Len(recs, 2)

	// Test GetByParam
	recs, err = repo.GetByParam(Params{
		{Field: "id", Value: 41},
	})
	r.NoError(err)
	r.Len(recs, 1)

	err = tx.Commit()
	r.NoError(err)

	// Test Update
	tx, err = database.Beginx()
	r.NoError(err)

	repo = New[User]("users", tx)
	err = repo.Update(41, &User{
		Id:        41,
		Email:     "john@smith.com",
		FirstName: "John",
		LastName: sql.NullString{
			String: "Smith",
			Valid:  true,
		},
	})
	r.NoError(err)

	err = tx.Commit()
	r.NoError(err)

	tx, err = database.Beginx()
	r.NoError(err)

	repo = New[User]("users", tx)
	rec, err = repo.Get(41)
	r.NoError(err)
	r.Equal("john@smith.com", rec.Email)
	r.Equal("John", rec.FirstName)
	r.Equal("Smith", rec.LastName.String)

	// Test Delete
	err = repo.Delete(40)
	r.NoError(err)

	err = repo.Delete(41)
	r.NoError(err)

	recs, err = repo.GetByParam(Params{
		{Field: "id", Operator: GreaterThan, Value: 1},
	})
	r.NoError(err)
	r.Len(recs, 0)
}
