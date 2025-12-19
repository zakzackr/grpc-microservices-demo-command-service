package presen

import (
	"github.com/zakzackr/grpc-microservices-demo-command-service/application"
	"github.com/zakzackr/grpc-microservices-demo-command-service/presen/adapter"
	"github.com/zakzackr/grpc-microservices-demo-command-service/presen/prepare"
	"github.com/zakzackr/grpc-microservices-demo-command-service/presen/server"
	"go.uber.org/fx"
)

var CommandDepend = fx.Options(
	application.SrvDepend, // アプリケーション層の依存定義
	fx.Provide( // プレゼンテーション層の依存定義
		adapter.NewCategoryAdapterImpl, // カテゴリ変換
		adapter.NewProductAdapterImpl,  // 商品変換
		server.NewCategoryServer,        // カテゴリサーバ
		server.NewProductServer,          // 商品サーバ
		prepare.NewCommandServer,        // gRPCサーバ
	),
	// メソッドの起動
	fx.Invoke(prepare.CommandServiceLifecycle),
)