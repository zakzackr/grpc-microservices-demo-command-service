package service

import (
	"context"

	"github.com/zakzackr/grpc-microservices-demo-command-service/domain/models/categories"
)

// カテゴリサービス
type CategoryService interface {
	// カテゴリを追加する
	Add(ctx context.Context, category *categories.Category) error
	// カテゴリを更新する
	Update(ctx context.Context, category *categories.Category) error
	// カテゴリを削除する
	Delete(ctx context.Context, category *categories.Category) error
}