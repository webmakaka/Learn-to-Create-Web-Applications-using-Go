package main

import (
	"fmt"
	"net/http"

	"./controllers"
	"./models"

	"github.com/gorilla/mux"
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

	must(err)

	defer us.Close()
	//us.DestructiveReset()
	us.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")

	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
