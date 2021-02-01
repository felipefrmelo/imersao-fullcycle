package model

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewPixKey(t *testing.T) {
	bank, _ := NewBank("123", "itau")
	account, _ := NewAccount(bank, "10291092091", "Felipe")
	pix, _ := NewPixKey("email", account, "Felipe")
	type args struct {
		kind    string
		account *Account
		key     string
	}
	tests := []struct {
		name    string
		args    args
		want    *PixKey
		wantErr bool
	}{
		{name: "Cria uma pixkey valida", want: pix, wantErr: false,
			args: args{kind: "email", account: account, key: "1212"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewPixKey(tt.args.kind, tt.args.account, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPixKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestModel_NewPixKey(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, err := NewAccount(bank, accountNumber, ownerName)

	kind := "email"
	key := "j@j.com"
	pixKey, err := NewPixKey(kind, account, key)

	require.NotEmpty(t, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Status, "active")

	kind = "cpf"
	_, err = NewPixKey(kind, account, key)
	require.Nil(t, err)

	_, err = NewPixKey("nome", account, key)
	require.NotNil(t, err)
}
