package main

import (
	"fmt"
	"log"
	"net"

	"github.com/k0kubun/pp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"todo.com/proto/greeting"
	pb "todo.com/proto/todo"
	"todo.com/services/todo/rpc"
)

const (
	RPC_PORT      = 4001
	GREETING_PORT = 4000
	GREETING_HOST = "172.18.0.3"
)

func main() {
	log.Print("Todo Server Starting")

	listenProt, err := net.Listen("tcp", fmt.Sprintf(":%d", RPC_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	conn, err := initConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	pp.Printf("%v\n", conn)

	grpc := grpc.NewServer()
	rpc := rpc.NewRPC(greeting.NewGreeterClient(conn))

	pb.RegisterGreeterServer(grpc, rpc)
	reflection.Register(grpc)

	if err := grpc.Serve(listenProt); err != nil {
		log.Fatal(err)
	}
}

func initConnection() (*grpc.ClientConn, error) {
	// ① greeting-todo containerのIPを取得。
	// % make into-greeting
	// # ifconfig
	// eth0      Link encap:Ethernet  HWaddr 02:42:AC:12:00:03
	//           inet addr:172.18.0.3  Bcast:172.18.255.255  Mask:255.255.0.0
	//           UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
	//           RX packets:42 errors:0 dropped:0 overruns:0 frame:0
	//           TX packets:23 errors:0 dropped:0 overruns:0 carrier:0
	//           collisions:0 txqueuelen:0
	//           RX bytes:3212 (3.1 KiB)  TX bytes:1825 (1.7 KiB)
	//
	// ② コンテナ間の疎通ができるか確認する
	// % make into-todo
	// # ping -c 3 172.18.0.3
	// PING 172.18.0.3 (172.18.0.3): 56 data bytes
	// 64 bytes from 172.18.0.3: seq=0 ttl=64 time=0.788 ms
	// 64 bytes from 172.18.0.3: seq=1 ttl=64 time=0.547 ms
	// 64 bytes from 172.18.0.3: seq=2 ttl=64 time=1.264 ms
	//
	// --- 172.18.0.3 ping statistics ---
	// 3 packets transmitted, 3 packets received, 0% packet loss
	// round-trip min/avg/max = 0.547/0.866/1.264 ms
	//
	// ③ コンテナ間で疎通が確認できたら、rpc client側のhostを「172.18.0.3」に設定
	// 		もしくは、container nameを設定すると勝手に名前解決してくれる
	// ④ なので、以下のどちらでもok。しかしIPが変更されるので、名前解決を推奨
	// connTO := fmt.Sprintf("172.18.0.3:%d", 4000)
	connTO := fmt.Sprintf("todo-greeting:%d", 4000)
	conn, err := grpc.Dial(
		connTO,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return conn, nil
}
