package userdom

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumberMaxLength(t *testing.T) {
	require := require.New(t)

	Number, myError := NewNumber(strings.Repeat("1", maxNumberLength+1))
	require.ErrorIs(myError, ErrNumberMaxLength)
	require.Len(Number, 0)
}

func TestValidNumber(t *testing.T) {
	require := require.New(t)

	number, myError := NewNumber(strings.Repeat("1", maxNumberLength-1))
	require.Nil(myError)
	require.NotEmpty(number)
}

func TestEqualNumber(t *testing.T) {
	require := require.New(t)

	number, myError := NewNumber(strings.Repeat("1", maxNumberLength-1))
	require.Nil(myError)

	number2, myError2 := NewNumber(strings.Repeat("1", maxNumberLength-1))
	require.Nil(myError2)

	require.True(number.Equals(number2))

}

func TestNotEqualNumber(t *testing.T) {
	require := require.New(t)

	number, myError := NewNumber(strings.Repeat("1", maxNumberLength-1))
	require.Nil(myError)

	number2, myError2 := NewNumber(strings.Repeat("2", maxNumberLength-1))
	require.Nil(myError2)

	number3, myError3 := NewNumber(strings.Repeat("1", maxNumberLength))
	require.Nil(myError3)

	require.False(number.Equals(number3))
	require.False(number.Equals(number2))

}
