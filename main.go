package main

import (
	"todo/postgres"
)

func main(){
	DB := postgres.New(&pg.Options{
		User: "postgres",Password: "postgres",Database: "todo_dev",
	})
	defer DB.Close()
}