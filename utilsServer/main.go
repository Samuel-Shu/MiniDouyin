package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "utilsServer/rpcPb"
	"utilsServer/rpcServer"
)

func main()  {
	port := ":8088"
	listen, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("listen:%v",port)
	s := grpc.NewServer()
	pb.RegisterUtilsServerServer(s,&rpcServer.UtilsServerServer{})
	err1 := s.Serve(listen)
	if err1 != nil {
		return
	}
}

