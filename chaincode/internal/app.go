package internal

import (
	"context"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
	"github.com/ilau020203/abac-hyperledger/internal/handler"
	"go.uber.org/fx"
)

type App struct {
	chaincode *contractapi.ContractChaincode
}

func NewApp(lc fx.Lifecycle, chaincode *handler.SimpleChaincode) *App {
	srv := &App{}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			var err error
			srv.chaincode, err = contractapi.NewChaincode(chaincode)
			if err != nil {
				return fmt.Errorf("error creating abac chaincode: %v", err)
			}

			if err = srv.chaincode.Start(); err != nil {
				return fmt.Errorf("error starting abac chaincode: %v", err)
			}

			return nil
		},
	})

	return srv
}
