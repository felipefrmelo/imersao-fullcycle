package repository

import (
	"os"
	"testing"

	"github.com/felipefrmelo/imersao-fullcycle/domain/model"
	"github.com/felipefrmelo/imersao-fullcycle/infrastructure/db"
	"github.com/stretchr/testify/require"
)

func TestPixKeyRepositoryDb_FindAccount(t *testing.T) {
	database := db.ConnectDB(os.Getenv("env"))
	var a model.PixKeyRepositoryInterface = PixKeyRepositoryDb{Db: database}
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, _ := model.NewAccount(bank, accountNumber, ownerName)
	err := a.AddAccount(account)
	require.Nil(t, err)
	accountFromBank, err := a.FindAccount(account.ID)
	require.Equal(t, accountFromBank.BankID, bank.ID)
	require.Equal(t, accountFromBank.ID, account.ID)
	require.NotSame(t, *accountFromBank, *account)
}
