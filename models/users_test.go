package models

import (
	"fmt"
	"testing"
	"time"
)

func testingUserService() (*UserService, error) {
	const (
		host     = "ec2-23-23-184-76.compute-1.amazonaws.com"
		port     = 5432
		user     = "hrgcmhzjkgllyf"
		password = "f867d132e78e27e50a27d0b7522dbf3f44dc835c903eb3040d74ecd5daf5c633"
		dbname   = "d61hvpjfrp6em7"
		sslmode  = "require"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	us, err := NewUserService(psqlInfo)
	if err != nil {
		return nil, err
	}

	us.db.LogMode(false)
	us.DestructiveReset()
	return us, nil

}

func TestCreateUser(t *testing.T) {
	us, err := testingUserService()

	if err != nil {
		t.Fatal(err)
	}

	user := User{
		Name:  "Michael Scott",
		Email: "michae@example.com",
	}

	err = us.Create(&user)
	if err != nil {
		t.Fatal(err)
	}

	if user.ID == 0 {
		t.Errorf("Excpected ID > 0. Reeived %d", user.ID)
	}

	if time.Since(user.CreatedAt) > time.Duration(5*time.Second) {
		t.Errorf("Expected CreatedAt to be recent. Received %d", user.ID)
	}
}
