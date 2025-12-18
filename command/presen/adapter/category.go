package adapter

import (
	"github.com/zakzackr/grpc-microservices-demo-command-service/domain/models/categories"
	"github.com/zakzackr/grpc-microservices-demo-pb/pb"
)

// データ変換インターフェース
type CategoryAdapter interface {
	// CategoryUpParamをCategoryエンティティに変換
	ToEntity(param *pb.CategoryUpParam) (*categories.Category, error) 
	// CategoryエンティティをCategoryUpResultに変換
	ToResult(result any)  *pb.CategoryUpResult
}