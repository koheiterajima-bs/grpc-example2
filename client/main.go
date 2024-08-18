package main

import (
	"context"
	"grpc-example2/pb"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// NewClientでgRPCサーバーに接続
	conn, err := grpc.NewClient("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 関数が終了する際に接続を閉じる
	defer conn.Close()

	// Greeterサービスのクライアントインスタンスを作成
	c := pb.NewGreeterClient(conn)

	// タイムアウト付きのコンテキストを設定する
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 関数が終了する際にタイムアウトをキャンセルする
	defer cancel()

	// SayHelloメソッドを呼び出し、HelloRequestをサーバーに送信する
	r, err := c.SayHello(ctx, &pb.HelloRequest{Requestname: "Terajima"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// レスポンスメッセージからMessageフィールドを取得する
	log.Printf("Greeting: %s", r.GetResponseoccupation())
}
