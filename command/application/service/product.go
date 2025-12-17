package service

import (
	"context"

	"github.com/zakzackr/grpc-microservices-demo-command-service/domain/models/products"
)

// 商品サービス
type ProductService interface {
	// 商品を追加する
	Add(ctx context.Context, product *products.Product) error
	// 商品を更新する
	Update(ctx context.Context, product *products.Product) error
	// 商品を削除する
	Delete(ctx context.Context, product *products.Product) error
}