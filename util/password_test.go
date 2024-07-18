package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(9)

	hashedPassword, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CompareHashPassword(password, hashedPassword)

	require.NoError(t, err)

	wrongPassword := RandomString(6)

	err = CompareHashPassword(wrongPassword, hashedPassword)

	require.Error(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword2, hashedPassword)

}
