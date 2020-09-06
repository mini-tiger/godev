package main

import (
	"context"
	"log"
	"net"
	"path/filepath"
	"runtime"

	pb "godev/mymodels/grpc/CutstomTLS/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"path"
)

type SearchService struct {
	auth *Auth
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	if err := s.auth.Check(ctx); err != nil {
		return nil, err
	}
	log.Printf("Recv Client Message:%s\n", r.GetRequest())
	return &pb.SearchResponse{Response: r.GetRequest() + " Token Server"}, nil
}

const PORT = "5003"

func main() {
	_, file, _, _ := runtime.Caller(0)
	Basedir := filepath.Dir(filepath.Dir(file))
	devDir := filepath.Join(Basedir, "perm")
	certFile := path.Join(devDir, "server1.pem")
	keyFile := path.Join(devDir, "server1.key")

	c, err := credentials.NewServerTLSFromFile(certFile, keyFile)

	if err != nil {
		log.Fatalf("tlsServer.GetTLSCredentials err: %v", err)
	}
	//自定义拦截器实现
	server := grpc.NewServer(grpc.Creds(c), grpc.UnaryInterceptor(TokenInterceptor))

	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", "0.0.0.0:"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	log.Printf("Start Listen PORT:%s\n", PORT)
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

//自定义拦截器实现
func TokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	//通过metadata
	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return nil, status.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	//fmt.Println(md)
	var appKey string
	var appSecret string

	if key, ok := md["app_key"]; ok {
		appKey = key[0]
	}

	if secret, ok := md["app_secret"]; ok {
		appSecret = secret[0]
	}

	if appKey != "eddycjy" || appSecret != "20181005" {
		return nil, status.Errorf(codes.Unauthenticated, "Token 不合法")
	}

	//通过token验证，继续处理请求
	return handler(ctx, req)
}
