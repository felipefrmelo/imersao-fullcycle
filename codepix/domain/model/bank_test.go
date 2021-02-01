package model

import (
	"testing"
)

func TestNewBank(t *testing.T) {

	bank, err := NewBank("123", "itau")

	if err != nil {
		t.Errorf("Erro ao criar um banco valido")
	}

	if bank.Code != "123" && bank.Name != "itau" {
		t.Errorf("Erro nos atributtos")
	}

	_, err = NewBank("", "itau")

	if err == nil {
		t.Errorf("Banco criado com atrbuttos vazio")
	}

}

func TestBank_isValid(t *testing.T) {
	tests := []struct {
		name    string
		bank    *Bank
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bank.isValid(); (err != nil) != tt.wantErr {
				t.Errorf("Bank.isValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
