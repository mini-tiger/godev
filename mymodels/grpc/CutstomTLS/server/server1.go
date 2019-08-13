package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	pb "godev/mymodels/grpc/CutstomTLS/proto"
	"path"
	"google.golang.org/grpc/credentials"
)

type SearchService struct {
	auth *Auth
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	if err := s.auth.Check(ctx); err != nil {
		return nil, err
	}
	log.Printf("Recv Client Message:%s\n",r.GetRequest())
	return &pb.SearchResponse{Response: r.GetRequest() + " Token Server"}, nil
}

const PORT = "5002"

func main() {
	devDir:="/home/go/src/godev/mymodels/grpc/CutstomTLS/perm"
	certFile := path.Join(devDir,"server1.pem")
	keyFile := path.Join(devDir,"server1.key")

	c, err := credentials.NewServerTLSFromFile(certFile,keyFile)

	if err != nil {
		log.Fatalf("tlsServer.GetTLSCredentials err: %v", err)
	}

	server := grpc.NewServer(grpc.Creds(c))

	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", "0.0.0.0:"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	log.Printf("Start Listen PORT:%s\n",PORT)
	server.Serve(lis)

}


type Auth struct {
	appKey    string
	appSecret string
}

func (a *Auth) Check(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext err")
	}

	var (
		appKey    string
		appSecret string
	)
	if value, ok := md["app_key"]; ok {
		appKey = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}

	if appKey != a.GetAppKey() || appSecret != a.GetAppSecret() {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil
}

func (a *Auth) GetAppKey() string {
	return "eddycjy"
}

func (a *Auth) GetAppSecret() string {
	return "20181005"
}