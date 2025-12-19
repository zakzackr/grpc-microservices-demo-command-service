package prepare

import (
	"github.com/zakzackr/grpc-microservices-demo-pb/pb"
	"google.golang.org/grpc"
)

// gRPCサーバ
type CommandServer struct {
	Server *grpc.Server
}

// コンストラクタ
func NewCommandServer(category pb.CategoryCommandServer, product pb.ProductCommandServer) *CommandServer {
	// gRPCサーバを生成
	server := grpc.NewServer()
	// CategoryCommandServerを登録
	pb.RegisterCategoryCommandServer(server, category)
	// ProductCommandServerを登録
	pb.RegisterProductCommandServer(server, product)
	return &CommandServer{Server: server}
}