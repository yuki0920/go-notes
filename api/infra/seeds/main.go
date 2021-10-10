package main

import (
	"log"
	"os"
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/injector"

	_ "github.com/go-sql-driver/mysql" // MySQLのドライバーを使う
)

func main() {
	userRepo := injector.InjectUserRepository()

	user := &model.User{}

	user.Name = os.Args[1]

	password := os.Args[2]
	if err := user.SetPassword(password); err != nil {
		log.Fatalf("Set password failed %v", err)
	}

	if err := userRepo.Create(user); err != nil {
		log.Fatalf("User create failed %v", err)
	}

	log.Println("user registered successfully")
}
