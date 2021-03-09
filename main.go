package main

import (
	"fmt"
	"go-to-do/app/controllers"
	"go-to-do/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
