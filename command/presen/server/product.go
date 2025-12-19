package server

import (
	"context"

	"github.com/zakzackr/grpc-microservices-demo-command-service/application/service"
	"github.com/zakzackr/grpc-microservices-demo-command-service/presen/adapter"
	"github.com/zakzackr/grpc-microservices-demo-pb/pb"
)

type ProductServer struct {
	adapter adapter.ProductAdapter // データ変換
	service service.ProductService
	pb.UnimplementedProductCommandServer // embed
}

// コンストラクタ
func NewProductServer(adapter adapter.ProductAdapter, service service.ProductService) pb.ProductCommandServer {
	return &ProductServer{adapter: adapter, service: service}
}

// 商品の追加
// pb.UnimplementedProductCommandServerのCreateメソッドのオーバーライド
func (ins *ProductServer) Create(ctx context.Context, param *pb.ProductUpParam) (*pb.ProductUpResult, error) {
	if product, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		if err := ins.service.Add(ctx, product); err != nil {
			return ins.adapter.ToResult(err), nil
		}
		return ins.adapter.ToResult(product), nil
	}
}

// 商品の更新
// pb.UnimplementedProductCommandServerのUpdateメソッドのオーバーライド
func (ins *ProductServer) Update(ctx context.Context, param *pb.ProductUpParam) (*pb.ProductUpResult, error) {
	if product, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		if err := ins.service.Update(ctx, product); err != nil {
			return ins.adapter.ToResult(err), nil
		}
		return ins.adapter.ToResult(product), nil
	}
}

// 商品の削除
// pb.UnimplementedProductCommandServerのDeleteメソッドのオーバーライド
func (ins *ProductServer) Delete(ctx context.Context, param *pb.ProductUpParam) (*pb.ProductUpResult, error) {
	if product, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		if err := ins.service.Delete(ctx, product); err != nil {
			return ins.adapter.ToResult(err), nil
		}
		return ins.adapter.ToResult(product), nil
	}
}

