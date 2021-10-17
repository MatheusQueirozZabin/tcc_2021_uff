package uservo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDistrictMaxLength(t *testing.T) {
	require := require.New(t)

	district, myError := NewDistrict(strings.Repeat("a", MaxDistrictLength+1))
	require.ErrorIs(myError, ErrDistrictMaxLength)

	require.Len(district, 0)
}

func TestDistrictMinLength(t *testing.T) {
	require := require.New(t)

	district, myError := NewDistrict(strings.Repeat("a", MinDistrictLength-1))
	require.ErrorIs(myError, ErrDistrictMinLength)

	require.Len(district, 0)
}

func TestValidDistrict(t *testing.T) {
	require := require.New(t)

	district, myError := NewDistrict(strings.Repeat("a", MaxDistrictLength))

	require.Nil(myError)
	require.NotEmpty(district)
}

func TestEqualDistrict(t *testing.T) {
	require := require.New(t)

	district, myError := NewDistrict(strings.Repeat("a", MaxDistrictLength))
	require.Nil(myError)

	district2, myError2 := NewDistrict(strings.Repeat("a", MaxDistrictLength))
	require.Nil(myError2)

	require.True(district.Equals(district2))

}

func TestNotEqualDistrict(t *testing.T) {
	require := require.New(t)

	district, myError := NewDistrict(strings.Repeat("a", MaxDistrictLength))
	require.Nil(myError)

	district2, myError2 := NewDistrict(strings.Repeat("b", MaxDistrictLength))
	require.Nil(myError2)

	district3, myError3 := NewDistrict(strings.Repeat("a", MaxDistrictLength-1))
	require.Nil(myError3)

	require.False(district.Equals(district3))
	require.False(district.Equals(district2))

}
