package users

import (
	"encoding/json"
	"fmt"
	"log"
	d "ta/db"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	d.DB = sqlxDB

	u := NewUser(2, "oleg", "djur", "golang")

	u1, err := json.Marshal(u.Data)
	if err != nil {
		log.Fatal("error json marshal: ", err)
	}

	mock.ExpectExec(`INSERT INTO users`).WithArgs(u.ID, u1).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := u.Create(); err != nil {
		log.Fatal(err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRead(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	d.DB = sqlxDB

	u := NewUser(1, "oleg", "djur", "golang")
	u1, _ := json.Marshal(u.Data)

	rows := sqlmock.NewRows([]string{"id", "data"}).AddRow(u.ID, u1)
	mock.ExpectQuery("SELECT").WithArgs(u.ID).WillReturnRows(rows)

	if err = u.Read(); err != nil {
		log.Printf("error read %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		fmt.Println("got error:", err)
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	d.DB = sqlxDB

	u := NewUser(1, "oleg", "djur", "golang")

	u1, _ := json.Marshal(u.Data)

	mock.ExpectExec(`UPDATE users`).WithArgs(u1, u.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	if err = u.Update(); err != nil {
		log.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		log.Fatal("got error:", err)
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	d.DB = sqlxDB

	u := NewUser(1, "oleg", "djur", "golang")

	mock.ExpectExec(`DELETE`).WithArgs(u.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	if err = u.Delete(); err != nil {
		log.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		log.Fatal("got error:", err)
	}
}
