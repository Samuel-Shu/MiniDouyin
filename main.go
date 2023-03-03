package main

import (
	"MiniDouyin/db"
	"MiniDouyin/router"
)

func main() {
	db.InitDb()
	db.InitRdb()
	router.InitRouter()
}
