package main

import (
	"context"
	"grpc-example2/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

// gRPCサーバー実装で使用されるGoの構造体
type server struct {
	pb.UnimplementedGreeterServer
}

// reqの部分をリクエストで埋め、返す(これはいつ使われる？)
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Responseoccupation: "Engineer " + req.Requestname}, nil
}

func main() {
	// TCPリスナーをポート50052で作成する
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 新しいgRPCサーバーインスタンスを作成する
	s := grpc.NewServer()
	// Greeterサービスをサーバーインスタンスに登録する
	pb.RegisterGreeterServer(s, &server{})

	// サーバーがリスニングしているアドレスをログに出力
	log.Printf("server listening at %v", lis.Addr())
	// サーバーを起動し、リスナーで受け取ったリクエストを処理する
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
