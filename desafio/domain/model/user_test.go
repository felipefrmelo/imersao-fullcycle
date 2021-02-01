package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {

	felipe, err := NewUser("felipe", "felipe@gmail.com")
	require.Equal(t, felipe.Name, "felipe")
	require.Equal(t, felipe.Email, "felipe@gmail.com")
	require.Nil(t, err)

	usuarioComEmailInvalido, err := NewUser("felipe", "felipegmail.com")
	require.Nil(t, usuarioComEmailInvalido)
	require.Error(t, err)
}
