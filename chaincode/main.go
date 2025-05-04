package main

import (
	"github.com/ilau020203/abac-hyperledger/internal"
	"github.com/ilau020203/abac-hyperledger/internal/businesslogic"
	"github.com/ilau020203/abac-hyperledger/internal/handler"
	"github.com/ilau020203/abac-hyperledger/internal/repository"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(handler.NewService),
		fx.Provide(businesslogic.NewService),
		fx.Provide(repository.NewService),
		fx.Provide(internal.NewApp),

		fx.Invoke(func(*internal.App) {}),
	).Run()
}
