package main

import (
	"github.com/zakzackr/grpc-microservices-demo-command-service/presen"
	"go.uber.org/fx"
)

func main() {
	// fxを起動する
	fx.New(
		presen.CommandDepend, // 依存性を定義する
	).Run()
}