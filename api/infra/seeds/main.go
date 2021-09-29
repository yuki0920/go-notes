package main

import (
	"log"
	"os"
	database "yuki0920/go-blog/db"
	"yuki0920/go-blog/model"
	"yuki0920/go-blog/repository"

	_ "github.com/go-sql-driver/mysql" // MySQLのドライバーを使う
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("db connection failed %v", err)
	} else {
		log.Println("db connection established")
	}

	repository.SetDB(db)

	user := &model.User{}

	user.Name = os.Args[1]

	password := os.Args[2]
	if err := user.SetPassword(password); err != nil {
		log.Fatalf("Set password failed %v", err)
	}

	if _, err := repository.UserCreate(user); err != nil {
		log.Fatalf("User create failed %v", err)
	}

	log.Println("user registered successfully")
}
