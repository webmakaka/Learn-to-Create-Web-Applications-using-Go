package main

import (
	"fmt"

	"../models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "ec2-23-23-184-76.compute-1.amazonaws.com"
	port     = 5432
	user     = "hrgcmhzjkgllyf"
	password = "f867d132e78e27e50a27d0b7522dbf3f44dc835c903eb3040d74ecd5daf5c633"
	dbname   = "d61hvpjfrp6em7"
	sslmode  = "require"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	us, err := models.NewUserService(psqlInfo)

	if err != nil {
		panic(err)
	}

	defer us.Close()

	// Delete Database
	// 	us.DestructiveReset()

	// user := models.User{
	// 	Name:  "Michael Scott",
	// 	Email: "michae@example.com",
	// }

	// if err := us.Create(&user); err != nil {
	// 	panic(err)
	// }

	user, err := us.ByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
