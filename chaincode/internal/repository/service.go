package repository

import (
	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

type Service struct {
	simpleRepository *SimpleRepository
	ctx              contractapi.TransactionContextInterface
}

func NewService() *Service {
	return &Service{
		simpleRepository: NewSimpleRepository(nil),
		ctx:              nil,
	}
}

func (s *Service) SimpleRepo() *SimpleRepository {
	return s.simpleRepository
}

func (s *Service) GetContext() contractapi.TransactionContextInterface {
	return s.ctx
}

func (s *Service) SetContext(ctx contractapi.TransactionContextInterface) {
	s.ctx = ctx
	s.simpleRepository.SetContext(ctx)
}
