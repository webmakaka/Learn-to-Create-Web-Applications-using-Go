package main


type Dog struct {

}




//////////////////////////////////////
//////////////////////////////////////

// import (
// 	"fmt"

// 	"../models"
// )

// const (
// 	host     = "ec2-23-23-184-76.compute-1.amazonaws.com"
// 	port     = 5432
// 	user     = "hrgcmhzjkgllyf"
// 	password = "f867d132e78e27e50a27d0b7522dbf3f44dc835c903eb3040d74ecd5daf5c633"
// 	dbname   = "d61hvpjfrp6em7"
// 	sslmode  = "require"
// )

// func main() {
// 	// toHash := []byte("this is my string to hash")
// 	// h := hmac.New(sha256.New, []byte("my-secret-key"))
// 	// h.Write(toHash)
// 	// b := h.Sum(nil)

// 	// fmt.Println(base64.URLEncoding.EncodeToString(b))

// 	// hmac := hash.NewHMAC("my-secret-key")
// 	// fmt.Println(hmac.Hash("this is my string to hash"))

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

// 	us, err := models.NewUserService(psqlInfo)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer us.Close()
// 	us.DestructiveReset()
// 	// us.AutoMigrate()

// 	user := models.User{
// 		Name:     "Joh Calhoun",
// 		Email:    "jon@calhoun.io",
// 		Password: "jon",
// 		Remember: "abc123",
// 	}

// 	err = us.Create(&user)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("%+v\n", user)

// 	user2, err := us.ByRemember("abc123")

// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%+v\n", *user2)
// }

//////////////////////////////////////
//////////////////////////////////////

// package main

// import (
// 	"fmt"

// 	"../rand"
// )

// func main() {
// 	fmt.Println(rand.String(10))
// 	fmt.Println(rand.RememberToken())
// }

//////////////////////////////////////
//////////////////////////////////////

// package main

// import (
// 	"fmt"

// 	"../models"

// 	_ "github.com/jinzhu/gorm/dialects/postgres"
// )

// const (
// 	host     = "ec2-23-23-184-76.compute-1.amazonaws.com"
// 	port     = 5432
// 	user     = "hrgcmhzjkgllyf"
// 	password = "f867d132e78e27e50a27d0b7522dbf3f44dc835c903eb3040d74ecd5daf5c633"
// 	dbname   = "d61hvpjfrp6em7"
// 	sslmode  = "require"
// )

// func main() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

// 	us, err := models.NewUserService(psqlInfo)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer us.Close()

// 	// Delete Database
// 	us.DestructiveReset()

// 	user := models.User{
// 		Name:  "Michael Scott",
// 		Email: "michael@example.com",
// 	}

// 	if err := us.Create(&user); err != nil {
// 		panic(err)
// 	}

// 	if err := us.Delete(user.ID); err != nil {
// 		panic(err)
// 	}

// 	userById, err := us.ByID(user.ID)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(userById)
// }
