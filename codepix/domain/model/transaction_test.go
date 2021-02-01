package model

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"
)

func TestNewTransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, _ := NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "abcdestination"
	ownerName = "Mariana"
	accountDestination, _ := NewAccount(bank, accountNumberDestination, ownerName)

	kind := "email"
	key := "j@j.com"
	pixKey, _ := NewPixKey(kind, accountDestination, key)

	require.NotEqual(t, account.ID, accountDestination.ID)

	amount := 3.10
	transaction, err := NewTransaction(account, amount, pixKey, "My description")
	//
	require.Nil(t, err)
	require.NotNil(t, uuid.FromStringOrNil(transaction.ID))
	require.Equal(t, transaction.Amount, amount)
	require.Equal(t, transaction.Status, TransactionPending)
	require.Equal(t, transaction.Description, "My description")
	require.Empty(t, transaction.CancelDescription)

	pixKeySameAccount, err := NewPixKey(kind, account, key)

	_, err = NewTransaction(account, amount, pixKeySameAccount, "My description")
	require.NotNil(t, err)

	_, err = NewTransaction(account, 0, pixKey, "My description")
	require.NotNil(t, err)

}

func TestModel_ChangeStatusOfATransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, _ := NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "abcdestination"
	ownerName = "Mariana"
	accountDestination, _ := NewAccount(bank, accountNumberDestination, ownerName)

	kind := "email"
	key := "j@j.com"
	pixKey, _ := NewPixKey(kind, accountDestination, key)

	amount := 3.10
	transaction, _ := NewTransaction(account, amount, pixKey, "My description")

	transaction.Complete()
	require.Equal(t, transaction.Status, TransactionCompleted)

	transaction.Cancel("Error")
	require.Equal(t, transaction.Status, TransactionError)
	require.Equal(t, transaction.CancelDescription, "Error")

}
