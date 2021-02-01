package model

import (
	"testing"
)

func createAccount() *Account {
	bank, _ := NewBank("itau", "abc")
	acc, _ := NewAccount(bank, "123131", "felipe")
	return acc
}
func TestAccount_isValid(t *testing.T) {

	tests := []struct {
		name    string
		account *Account
		wantErr bool
	}{
		{name: "abc", account: createAccount(), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.account.isValid(); (err != nil) != tt.wantErr {
				t.Errorf("Account.isValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
