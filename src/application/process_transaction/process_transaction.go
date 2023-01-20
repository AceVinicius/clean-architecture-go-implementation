package process_transaction

import (
	"clean_architecture/src/domain"
)

type ProcessTransaction struct {
	Repository domain.TransactionRepository
}

func NewProcessTransaction(repository domain.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := domain.NewTransaction()

	transaction.Id = input.Id
	transaction.AccountId = input.AccountId
	transaction.Amount = input.Amount

	invalidTransaction := transaction.IsValid()

	if invalidTransaction != nil {
		return p.rejectTransaction(transaction, invalidTransaction)
	}

	return p.approveTransaction(transaction)
}

func (p *ProcessTransaction) approveTransaction(transaction *domain.Transaction) (TransactionDtoOutput, error) {
	transaction.Status = "approved"
	transaction.ErrorMessage = ""

	err := p.Repository.Insert(transaction.Id, transaction.AccountId, transaction.Amount, transaction.Status, transaction.ErrorMessage)

	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		Id:           transaction.Id,
		Status:       transaction.Status,
		ErrorMessage: transaction.ErrorMessage,
	}

	return output, nil
}

func (p *ProcessTransaction) rejectTransaction(transaction *domain.Transaction, invalidTransaction error) (TransactionDtoOutput, error) {
	transaction.Status = "rejected"
	transaction.ErrorMessage = invalidTransaction.Error()

	err := p.Repository.Insert(transaction.Id, transaction.AccountId, transaction.Amount, transaction.Status, transaction.ErrorMessage)

	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		Id:           transaction.Id,
		Status:       transaction.Status,
		ErrorMessage: transaction.ErrorMessage,
	}

	return output, nil
}
