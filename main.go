package main

import (
	"fmt"
	"go-to-do/app/controllers"
	"go-to-do/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
	/*
	user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)

	session, err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(session)

	valid, _ := session.CheckSession()
	fmt.Println(valid)

	 */
}
