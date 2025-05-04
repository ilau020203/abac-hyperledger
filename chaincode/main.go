package main

import (
	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
	"github.com/ilau020203/abac-hyperledger/internal"
	"github.com/ilau020203/abac-hyperledger/internal/businesslogic"
	"github.com/ilau020203/abac-hyperledger/internal/handler"
	"github.com/ilau020203/abac-hyperledger/internal/repository"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(func() contractapi.TransactionContextInterface {
			return nil
		}),
		fx.Provide(repository.NewService),
		fx.Provide(businesslogic.NewService),
		fx.Provide(func(bs *businesslogic.Service, rs *repository.Service) *handler.SimpleChaincode {
			return handler.NewAssetHandler(bs, rs)
		}),
		fx.Provide(internal.NewApp),

		fx.Invoke(func(*internal.App) {}),
	).Run()
}
