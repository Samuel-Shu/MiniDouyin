package main

import (
	"MiniDouyin/db"
	"MiniDouyin/router"
)

func main()  {
	db.InitDb()
	router.InitRouter()
}
