package main

import (
	"fmt"
	"golang-todo-app/app/controllers"
	"golang-todo-app/app/models"
)

func main(){

	fmt.Println(models.Db)

	controllers.StartMainServer()

}
