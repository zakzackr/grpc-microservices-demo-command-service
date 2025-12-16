package products

import (
	"context"
	"database/sql"
)

// 商品リポジトリインターフェイス
type ProductRepository interface {
	// 同名の商品の存在チェック
	Exists(ctx context.Context, tran *sql.Tx, product *Product) error
	// 新しい商品を作成する
	Create(ctx context.Context, tran *sql.Tx, product *Product) error
	// 商品を変更する
	UpdateById(ctx context.Context, tran *sql.Tx, product *Product) error
	// 商品を削除する
	DeleteById(ctx context.Context, tran *sql.Tx, product *Product) error
}