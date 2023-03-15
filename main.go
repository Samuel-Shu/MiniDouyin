package main

import (
	"MiniDouyin/db"
	"MiniDouyin/router"
	"MiniDouyin/rpc/rpcClient"
)

func main() {
	db.InitDb()
	db.InitRdb()
	rpcClient.InitRpcServer()
	router.InitRouter()
}
