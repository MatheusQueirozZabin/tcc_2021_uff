package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserNameMaxLength(t *testing.T) {
	require := require.New(t)

	userName, myError := NewUserName(strings.Repeat("a", MaxUserNameLength+1))
	require.ErrorIs(myError, ErrUserNameMaxLength)
	require.Len(userName, 0)
}

func TestUserNameMinLength(t *testing.T) {
	require := require.New(t)

	userName, myError := NewUserName(strings.Repeat("a", MinUserNameLength-1))
	require.ErrorIs(myError, ErrUserNameMinLength)
	require.Len(userName, 0)

}

func TestUserNameAlphanumeric(t *testing.T) {
	require := require.New(t)

	userName, myError := NewUserName(strings.Repeat("$", MaxUserNameLength))
	require.ErrorIs(myError, ErrUserNameInvalidCharacter)
	require.Len(userName, 0)

}

func TestValidUserName(t *testing.T) {
	require := require.New(t)

	userName, myError := NewUserName(strings.Repeat("a", MaxUserNameLength))
	require.Nil(myError)
	require.NotEmpty(userName)
}

func TestEqualUserName(t *testing.T) {
	require := require.New(t)

	username, myError := NewUserName(strings.Repeat("a", MaxUserNameLength))
	require.Nil(myError)
	username2, myError2 := NewUserName(strings.Repeat("a", MaxUserNameLength))
	require.Nil(myError2)

	require.True(username.Equals(username2))

}

func TestNotEqualUserName(t *testing.T) {
	require := require.New(t)

	username, myError := NewUserName(strings.Repeat("a", MaxUserNameLength))
	require.Nil(myError)

	username2, myError2 := NewUserName(strings.Repeat("b", MaxUserNameLength))
	require.Nil(myError2)

	username3, myError3 := NewUserName(strings.Repeat("a", MaxUserNameLength-1))
	require.Nil(myError3)

	require.False(username.Equals(username2))
	require.False(username.Equals(username3))

}
