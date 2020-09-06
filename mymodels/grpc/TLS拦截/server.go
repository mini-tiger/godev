package main

import (
	"context"
	"fmt"
	"godev/mymodels/grpc/TLS拦截/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
	"path/filepath"
	"runtime"
)

type MathManager struct {
}

func (mm *MathManager) AddMethod(ctx context.Context, request *message.RequestArgs) (response *message.Response, err error) {

	fmt.Println(" 服务端 Add方法 ")
	result := request.Args1 + request.Args2
	fmt.Println(" 计算结果是：", result)
	response = new(message.Response)
	response.Code = 1
	response.Message = "执行成功"
	return response, nil

}

func main() {
	_, file, _, _ := runtime.Caller(0)
	Basedir := filepath.Dir(filepath.Dir(file))
	//fmt.Println(filepath.Join(Basedir,"TLS拦截","keys","server.key"))
	//TLS认证
	creds, err := credentials.NewServerTLSFromFile(filepath.Join(Basedir, "TLS拦截", "keys", "server.pem"), filepath.Join(Basedir, "TLS拦截", "keys", "server.key"))
	if err != nil {
		grpclog.Fatal("加载在证书文件失败", err)
	}

	//实例化grpc server, 开启TLS认证
	server := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(TokenInterceptor))

	message.RegisterMathServiceServer(server, new(MathManager))

	lis, err := net.Listen("tcp", ":8093")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}

//自定义拦截器实现
func TokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	//通过metadata
	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return nil, status.Errorf(codes.Unauthenticated, "无Token认证信息")
	}

	var appKey string
	var appSecret string

	if key, ok := md["appid"]; ok {
		appKey = key[0]
	}

	if secret, ok := md["appkey"]; ok {
		appSecret = secret[0]
	}

	if appKey != "hello" || appSecret != "20190812" {
		return nil, status.Errorf(codes.Unauthenticated, "Token 不合法")
	}

	//通过token验证，继续处理请求
	return handler(ctx, req)
}
