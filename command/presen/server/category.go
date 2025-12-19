package server

import (
	"context"

	"github.com/zakzackr/grpc-microservices-demo-command-service/application/service"
	"github.com/zakzackr/grpc-microservices-demo-command-service/presen/adapter"
	"github.com/zakzackr/grpc-microservices-demo-pb/pb"
)

type CategoryServer struct {
	adapter adapter.CategoryAdapter  // データ変換
	service service.CategoryService
	pb.UnimplementedCategoryCommandServer  // embed
}

// コンストラクタ
func NewCategoryServer(adapter adapter.CategoryAdapter, service service.CategoryService) pb.CategoryCommandServer {
	return &CategoryServer{adapter: adapter, service: service}
}

// カテゴリの追加
// pb.UnimplementedCategoryCommandServerのCreateメソッドのオーバーライド
func (ins *CategoryServer) Create(ctx context.Context, param *pb.CategoryUpParam) (*pb.CategoryUpResult, error){
	// pb.CategoryUpParamをCategory Entityに変換
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		if err := ins.service.Add(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil
		}
		return ins.adapter.ToResult(category), nil
	}
}

// カテゴリの更新
// pb.UnimplementedCategoryCommandServerのUpdateメソッドのオーバーライド
func (ins *CategoryServer) Update(ctx context.Context, param *pb.CategoryUpParam) (*pb.CategoryUpResult, error) {
	// pb.CategoryUpParamをCategory Entityに変換
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		if err := ins.service.Update(ctx, category); err != nil {
			return  ins.adapter.ToResult(err), nil
		}
		return ins.adapter.ToResult(category), nil
	}
}

// カテゴリの削除
// pb.UnimplementedCategoryCommandServerのDeleteメソッドのオーバーライド
func (ins *CategoryServer) Delete(ctx context.Context, param *pb.CategoryUpParam) (*pb.CategoryUpResult, error) {
	// pb.CategoryUpParamをCategory Entityに変換
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		if err := ins.service.Delete(ctx, category); err != nil {
			return  ins.adapter.ToResult(err), nil
		}
		return ins.adapter.ToResult(category), nil
	}
}

