package handler

import (
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
	"github.com/ilau020203/abac-hyperledger/internal/businesslogic"
	"github.com/ilau020203/abac-hyperledger/internal/repository"
)

type SimpleChaincode struct {
	contractapi.Contract
	simpleService *businesslogic.SimpleService
	bService      *businesslogic.Service
	repoService   *repository.Service
}

func NewService(businessService *businesslogic.Service, repoService *repository.Service) *SimpleChaincode {
	return NewAssetHandler(businessService, repoService)
}

func NewAssetHandler(businessService *businesslogic.Service, repoService *repository.Service) *SimpleChaincode {
	return &SimpleChaincode{
		simpleService: businessService.Simple(),
		bService:      businessService,
		repoService:   repoService,
	}
}

func (h *SimpleChaincode) InitLedger(ctx contractapi.TransactionContextInterface, A string, AVal string, B string, BVal string) error {
	h.simpleService.Ctx = ctx
	h.repoService.SetContext(ctx)

	aValue, err := strconv.Atoi(AVal)
	if err != nil {
		return err
	}

	bValue, err := strconv.Atoi(BVal)
	if err != nil {
		return err
	}

	return h.simpleService.InitLedger(A, aValue, B, bValue)
}

func (h *SimpleChaincode) InvokeTransfer(ctx contractapi.TransactionContextInterface, A string, B string, X string) error {
	h.simpleService.Ctx = ctx
	h.repoService.SetContext(ctx)

	value, err := strconv.Atoi(X)
	if err != nil {
		return err
	}

	return h.simpleService.InvokeTransfer(A, B, value)
}

func (h *SimpleChaincode) Query(ctx contractapi.TransactionContextInterface, name string) (string, error) {
	h.simpleService.Ctx = ctx
	h.repoService.SetContext(ctx)

	value, err := h.simpleService.QueryAccount(name)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(value), nil
}
