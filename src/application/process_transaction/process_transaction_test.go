package process_transaction

import (
	mock_domain "clean_architecture/src/domain/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransactionWhenItsValid(t *testing.T) {
	input := TransactionDtoInput{
		Id:        "1",
		AccountId: "1",
		Amount:    666,
	}

	expectedOutput := TransactionDtoOutput{
		Id:           "1",
		Status:       "approved",
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_domain.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.Id, input.AccountId, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransactionWhenItsNotValid(t *testing.T) {
	input := TransactionDtoInput{
		Id:        "1",
		AccountId: "1",
		Amount:    1500,
	}

	expectedOutput := TransactionDtoOutput{
		Id:           "1",
		Status:       "rejected",
		ErrorMessage: "you don't have limit for this transaction",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_domain.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.Id, input.AccountId, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
