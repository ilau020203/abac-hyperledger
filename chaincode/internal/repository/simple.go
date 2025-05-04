package repository

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

type SimpleAccount struct {
	ID    string `json:"ID"`
	Value int    `json:"value"`
}

type SimpleRepository struct {
	Ctx contractapi.TransactionContextInterface
}

func NewSimpleRepository(ctx contractapi.TransactionContextInterface) *SimpleRepository {
	return &SimpleRepository{
		Ctx: ctx,
	}
}

func (r *SimpleRepository) CreateAccount(account SimpleAccount) error {
	if r.Ctx == nil {
		return fmt.Errorf("transaction context is nil")
	}
	accountJSON, err := json.Marshal(account)
	if err != nil {
		return err
	}
	return r.Ctx.GetStub().PutState(account.ID, accountJSON)
}

func (r *SimpleRepository) ReadAccount(id string) (*SimpleAccount, error) {
	if r.Ctx == nil {
		return nil, fmt.Errorf("transaction context is nil")
	}
	accountJSON, err := r.Ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if accountJSON == nil {
		return nil, fmt.Errorf("account %s does not exist", id)
	}

	var account SimpleAccount
	err = json.Unmarshal(accountJSON, &account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *SimpleRepository) UpdateAccount(account SimpleAccount) error {
	if r.Ctx == nil {
		return fmt.Errorf("transaction context is nil")
	}
	accountJSON, err := json.Marshal(account)
	if err != nil {
		return err
	}
	return r.Ctx.GetStub().PutState(account.ID, accountJSON)
}

func (r *SimpleRepository) DeleteAccount(id string) error {
	return r.Ctx.GetStub().DelState(id)
}

func (r *SimpleRepository) AccountExists(id string) (bool, error) {
	if r.Ctx == nil {
		return false, fmt.Errorf("transaction context is nil")
	}
	accountJSON, err := r.Ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return accountJSON != nil, nil
}

func (r *SimpleRepository) GetAllAccounts() ([]*SimpleAccount, error) {
	resultsIterator, err := r.Ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var accounts []*SimpleAccount
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var account SimpleAccount
		err = json.Unmarshal(queryResponse.Value, &account)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}
	return accounts, nil
}

func (r *SimpleRepository) SetContext(ctx contractapi.TransactionContextInterface) {
	r.Ctx = ctx
}
