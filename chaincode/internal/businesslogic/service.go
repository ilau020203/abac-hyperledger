package businesslogic

import (
	"encoding/base64"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
	"github.com/ilau020203/abac-hyperledger/internal/repository"
)

type SimpleService struct {
	Repository *repository.SimpleRepository
	Ctx        contractapi.TransactionContextInterface
}

func NewSimpleService(repo *repository.SimpleRepository, ctx contractapi.TransactionContextInterface) *SimpleService {
	return &SimpleService{
		Repository: repo,
		Ctx:        ctx,
	}
}

type Service struct {
	simpleService *SimpleService
	repoService   *repository.Service
	ctx           contractapi.TransactionContextInterface
}

func NewService(repoService *repository.Service, ctx contractapi.TransactionContextInterface) *Service {
	return &Service{
		simpleService: NewSimpleService(repoService.SimpleRepo(), ctx),
		repoService:   repoService,
		ctx:           ctx,
	}
}

func (s *Service) Simple() *SimpleService {
	return s.simpleService
}

func (s *SimpleService) GetSubmittingClientIdentity() (string, error) {
	b64ID, err := s.Ctx.GetClientIdentity().GetID()
	if err != nil {
		return "", fmt.Errorf("failed to read clientID: %v", err)
	}
	decodeID, err := base64.StdEncoding.DecodeString(b64ID)
	if err != nil {
		return "", fmt.Errorf("failed to base64 decode clientID: %v", err)
	}
	return string(decodeID), nil
}

func (s *SimpleService) InitLedger(A string, AVal int, B string, BVal int) error {
	s.Repository.SetContext(s.Ctx)

	err := s.Ctx.GetClientIdentity().AssertAttributeValue("abac.init", "true")
	if err != nil {
		return fmt.Errorf("submitting client not authorized to initialize ledger, does not have abac.init=true attribute")
	}

	existsA, err := s.Repository.AccountExists(A)
	if err != nil {
		return err
	}
	if existsA {
		return fmt.Errorf("account %s already exists", A)
	}

	existsB, err := s.Repository.AccountExists(B)
	if err != nil {
		return err
	}
	if existsB {
		return fmt.Errorf("account %s already exists", B)
	}

	accountA := repository.SimpleAccount{
		ID:    A,
		Value: AVal,
	}
	err = s.Repository.CreateAccount(accountA)
	if err != nil {
		return fmt.Errorf("failed to create account A: %v", err)
	}

	accountB := repository.SimpleAccount{
		ID:    B,
		Value: BVal,
	}
	err = s.Repository.CreateAccount(accountB)
	if err != nil {
		return fmt.Errorf("failed to create account B: %v", err)
	}

	return nil
}

func (s *SimpleService) InvokeTransfer(A string, B string, X int) error {
	s.Repository.SetContext(s.Ctx)

	accountA, err := s.Repository.ReadAccount(A)
	if err != nil {
		return fmt.Errorf("failed to get state for %s: %v", A, err)
	}
	if accountA == nil {
		return fmt.Errorf("account %s not found", A)
	}

	accountB, err := s.Repository.ReadAccount(B)
	if err != nil {
		return fmt.Errorf("failed to get state for %s: %v", B, err)
	}
	if accountB == nil {
		return fmt.Errorf("account %s not found", B)
	}

	accountA.Value = accountA.Value - X
	accountB.Value = accountB.Value + X

	err = s.Repository.UpdateAccount(*accountA)
	if err != nil {
		return fmt.Errorf("failed to update account %s: %v", A, err)
	}

	err = s.Repository.UpdateAccount(*accountB)
	if err != nil {
		return fmt.Errorf("failed to update account %s: %v", B, err)
	}

	return nil
}

func (s *SimpleService) QueryAccount(id string) (int, error) {
	s.Repository.SetContext(s.Ctx)

	account, err := s.Repository.ReadAccount(id)
	if err != nil {
		return 0, fmt.Errorf("failed to get account %s: %v", id, err)
	}

	return account.Value, nil
}
