package main

import (
	"fmt"
	"golang-todo-app/config"
	"log"
)

func main(){
	fmt.Println("Called In main")
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")
}
