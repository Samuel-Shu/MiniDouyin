package rpcClient

import (
	"MiniDouyin/config"
	pb "MiniDouyin/rpc/rpcpb"
	"google.golang.org/grpc"
	"log"
)

var Client pb.UtilsServerClient

func InitRpcServer() {
	conn, err := grpc.Dial(config.RpcDNS, grpc.WithInsecure())
	if err != nil {
		log.Fatal("rpc dial fail", err)
	}
	defer conn.Close()
	Client = pb.NewUtilsServerClient(conn)
}
