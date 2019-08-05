package main

import (
	"context"
	"log"

	"google.golang.org/grpc"


	pb "godev/mymodels/grpc/CutstomTLS/proto"
	"path"
	"google.golang.org/grpc/credentials"
)

const PORT = "5002"

type Auth struct {
	AppKey    string
	AppSecret string
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"app_key": a.AppKey, "app_secret": a.AppSecret}, nil
}

func (a *Auth) RequireTransportSecurity() bool {
	return true
}

func main() {
	devDir:="/home/go/src/godev/mymodels/grpc/CutstomTLS/perm"
	certFile := path.Join(devDir,"server1.pem")


	c, err := credentials.NewClientTLSFromFile(certFile,
		"tj-test") // todo 认证时候的名字


	if err != nil {
		log.Fatalf("tlsClient.GetTLSCredentials err: %v", err)
	}

	auth := Auth{
		AppKey:    "eddycjy",
		AppSecret: "20181005",
	}
	conn, err := grpc.Dial("127.0.0.1:"+PORT, grpc.WithTransportCredentials(c), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	// todo 调用服务端search方法
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "tj gRPC",
	})

	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())
}
