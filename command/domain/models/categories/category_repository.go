package categories

import (
	"context"
	"database/sql"
)

// カテゴリリポジトリインターフェイス
type CategoryRepository interface {
	// 同名の商品カテゴリの存在チェック
	Exists(ctx context.Context, tran *sql.Tx, category *Category) error
	// 新しい商品カテゴリを作成する
	Create(ctx context.Context, tran *sql.Tx, category *Category) error
	// 商品カテゴリを変更する
	UpdateById(ctx context.Context, tran *sql.Tx, category *Category) error
	// 商品カテゴリを削除する
	DeleteById(ctx context.Context, tran *sql.Tx, category *Category) error
}