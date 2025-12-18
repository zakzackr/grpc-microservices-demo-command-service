package adaptor

import (
	"github.com/zakzackr/grpc-microservices-demo-command-service/domain/models/products"
	"github.com/zakzackr/grpc-microservices-demo-pb/pb"
)

type ProductAdaptor interface {
	// ProductUpParamをProductエンティティに変換
	ToEntity(param *pb.ProductUpParam) (*products.Product, error)
	// ProductエンティティをProductUpResultに変換
	ToResult(result any) *pb.ProductUpResult
}