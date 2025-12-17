package application

import (
	"github.com/zakzackr/grpc-microservices-demo-command-service/application/impl"
	"github.com/zakzackr/grpc-microservices-demo-command-service/infra/sqlboiler"
	"go.uber.org/fx"
)

// アプリケーション層の依存定義
var SrvDepend = fx.Options(
	sqlboiler.RepDepend, // SQLBoilderを利用したリポジトリインターフェイス実装
	fx.Provide(
		// サービスインターフェイス実装のコンストラクタ
		impl.NewCategoryServiceImpl,
		impl.NewProductServiceImpl,
	),
)