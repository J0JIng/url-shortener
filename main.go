package main

import(
	"learn-go/handlers"
	"learn-go/database"
)

func main(){
	db := database.Connect()
	handlers.Handler(db)
}