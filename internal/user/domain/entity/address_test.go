package userent_test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	userent "ifoodish-store/internal/user/domain/entity"
	uservo "ifoodish-store/internal/user/domain/valueobject"
	"ifoodish-store/pkg/resperr"

	"github.com/stretchr/testify/require"
)

var (
	validAddress = userent.Address{
		Street:     "Street ABCD",
		District:   "Espirito Santo",
		City:       "Jose dos Campos",
		State:      "Rio de Janeiro",
		Complement: "Complement",
		Number:     "11111",
		Zipcode:    "23970000",
		Latitude:   "-23.307577",
		Longitude:  "-44.754146",
	}
)

func TestAddressValid(t *testing.T) {
	require := require.New(t)
	address, err := userent.NewAddress(validAddress)
	require.Nil(err)
	require.NotEmpty(address)
}

func TestAddressInvalid(t *testing.T) {
	require := require.New(t)

	type testIterator struct {
		address userent.Address
		err     error
	}

	addresses := []testIterator{}

	example := validAddress
	example.Street = uservo.Street(strings.Repeat("a", uservo.MinStreetLength-1))
	addresses = append(addresses, testIterator{
		address: example,
		err:     uservo.ErrStreetMinLength,
	})

	example = validAddress
	example.Street = uservo.Street(strings.Repeat("a", uservo.MaxStreetLength+1))
	addresses = append(addresses, testIterator{
		address: example,
		err:     uservo.ErrStreetMaxLength,
	})

	// District
	example = validAddress
	example.District = uservo.District(strings.Repeat("a", uservo.MinDistrictLength-1))
	addresses = append(addresses, testIterator{
		address: example,
		err:     uservo.ErrDistrictMinLength,
	})

	example = validAddress
	example.District = uservo.District(strings.Repeat("a", uservo.MaxDistrictLength+1))
	addresses = append(addresses, testIterator{
		address: example,
		err:     uservo.ErrDistrictMaxLength,
	})
	// City
	example = validAddress
	example.City = uservo.City(strings.Repeat("a", uservo.MinCityLength-1))
	addresses = append(addresses, testIterator{
		address: example,
		err:     uservo.ErrCityMinLength,
	})

	example = validAddress
	example.City = uservo.City(strings.Repeat("a", uservo.MaxCityLength+1))
	addresses = append(addresses, testIterator{
		address: example,
		err:     uservo.ErrCityMaxLength,
	})
	// State
	example = validAddress
	example.State = uservo.State(strings.Repeat("a", uservo.MinStateLength-1))
	addresses = append(addresses, testIterator{
		address: example,
		err:     uservo.ErrStateMinLength,
	})

	example = validAddress
	example.State = uservo.State(strings.Repeat("a", uservo.MaxStateLength+1))
	addresses = append(addresses, testIterator{
		address: example,
		err:     uservo.ErrStateMaxLength,
	})
	// Complement
	example = validAddress
	example.Complement = uservo.Complement(strings.Repeat("a", uservo.MaxComplementLength+1))
	addresses = append(addresses, testIterator{
		address: example,
		err:     uservo.ErrComplementMaxLength,
	})
	// Number
	example = validAddress
	example.Number = uservo.AddressNumber(strings.Repeat("a", uservo.MinAddressNumberLength-1))
	addresses = append(addresses, testIterator{
		address: example,
		err:     uservo.ErrAddressNumberMinLength,
	})

	example = validAddress
	example.Number = uservo.AddressNumber(strings.Repeat("a", uservo.MaxAddressNumberLength+1))
	addresses = append(addresses, testIterator{
		address: example,
		err:     uservo.ErrAddressNumberMaxLength,
	})
	// Zipcode
	example = validAddress
	example.Zipcode = uservo.Zipcode(strings.Repeat("1", uservo.ZipcodeLength+1))
	addresses = append(addresses, testIterator{
		address: example,
		err:     uservo.ErrZipcodeLength,
	})
	example = validAddress
	example.Zipcode = uservo.Zipcode(strings.Repeat("a", uservo.ZipcodeLength))
	addresses = append(addresses, testIterator{
		address: example,
		err:     uservo.ErrZipcodeNotNumeric,
	})

	// Latitude
	example = validAddress
	example.Latitude = uservo.Latitude(strings.Repeat("a", 5))
	addresses = append(addresses, testIterator{
		address: example,
		err:     uservo.ErrLatitudeInvalidFormat,
	})

	// Longitude
	example = validAddress
	example.Longitude = uservo.Longitude(strings.Repeat("a", 5))
	addresses = append(addresses, testIterator{
		address: example,
		err:     uservo.ErrLongitudeInvalidFormat,
	})

	for _, it := range addresses {
		_, err := userent.NewAddress(it.address)
		require.ErrorIs(err, it.err)
	}

}

func TestRegisteredAddressValid(t *testing.T) {
	require := require.New(t)

	ex := userent.RegisteredAddress{}
	ex.Address = validAddress
	ex.ID = 50

	address, err := userent.NewRegisteredAddress(ex)
	require.Nil(err)
	require.NotEmpty(address)
}

func TestRegisteredAddressInvalidID(t *testing.T) {
	require := require.New(t)

	ex := userent.RegisteredAddress{}
	ex.Address = validAddress
	ex.ID = 0
	_, err := userent.NewRegisteredAddress(ex)
	require.ErrorIs(err, uservo.ErrInvalidAddressID)

	ex.ID = -10
	_, err = userent.NewRegisteredAddress(ex)
	require.ErrorIs(err, uservo.ErrInvalidAddressID)

	ex.City = uservo.City(strings.Repeat("a", uservo.MinCityLength-1))
	_, err = userent.NewRegisteredAddress(ex)
	require.ErrorIs(err, uservo.ErrCityMinLength)

}

func TestJSONUnmarshallingAddressSuccess(t *testing.T) {
	require := require.New(t)

	var address *userent.Address
	err := json.Unmarshal([]byte(`
	{
		"Street":     "Street ABCD",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.Nil(err)
	require.True(address.Street.Equals("Street ABCD"))
	require.True(address.District.Equals("District"))
	require.True(address.City.Equals("City"))
	require.True(address.State.Equals("State"))
	require.True(address.Complement.Equals("Complement"))
	require.True(address.Number.Equals("11111"))
	require.True(address.Zipcode.Equals("23970000"))
	require.True(address.Latitude.Equals("-23.307577"))
	require.True(address.Longitude.Equals("-44.754146"))
}

func TestJSONUnmarshallingAddressFail(t *testing.T) {
	require := require.New(t)
	var address *userent.Address

	// forçando teste do unmarshal
	err := address.UnmarshalJSON([]byte(`
		{
			"Street":     "Street ABCD",
			"District":   "District",
			"City":       "City",
			"State":      "State",
			"Complement": "Complement",
			"Number":     "11111",
			"Zipcode":    "23970000",
			"Latitude":   "-23.307577",
			"Longitude":  "-44.754146
		}
	`))
	require.NotNil(err)

	address = nil
	err = json.Unmarshal([]byte(`
	{
		"Street":     "",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrStreetMinLength)

}

//////////////////////////

func TestJSONUnmarshallingRegisteredAddressSuccess(t *testing.T) {
	require := require.New(t)

	var address userent.RegisteredAddress
	err := address.UnmarshalJSON([]byte(`
	{
		"id":         50,	
		"Street":     "Street ABCD",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`))
	require.Nil(err)
	require.True(address.ID.Equals(50))
	require.True(address.Street.Equals("Street ABCD"))
	require.True(address.District.Equals("District"))
	require.True(address.City.Equals("City"))
	require.True(address.State.Equals("State"))
	require.True(address.Complement.Equals("Complement"))
	require.True(address.Number.Equals("11111"))
	require.True(address.Zipcode.Equals("23970000"))
	require.True(address.Latitude.Equals("-23.307577"))
	require.True(address.Longitude.Equals("-44.754146"))
}

func TestJSONUnmarshallingRegisteredAddressFail(t *testing.T) {
	require := require.New(t)
	var address userent.RegisteredAddress

	// forçando teste do unmarshal
	err := address.UnmarshalJSON([]byte(`
		{
			"id":         50,	
			"Street":     "Street ABCD",
			"District":   "District",
			"City":       "City",
			"State":      "State",
			"Complement": "Complement",
			"Number":     "11111",
			"Zipcode":    "23970000",
			"Latitude":   "-23.307577",
			"Longitude":  "-44.754146
		}
	`))
	require.Equal(http.StatusInternalServerError, resperr.StatusCode(err))

	err = json.Unmarshal([]byte(`
	{
		"id":         -1,
		"Street":     "Street ABC",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrInvalidAddressID)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
		"Street":     "",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrStreetMinLength)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
		"Street":     "Street ABC",
		"District":   "",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrDistrictMinLength)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
		"Street":     "Street ABC",
		"District":   "District",
		"City":       "",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrCityMinLength)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
		"Street":     "Street ABC",
		"District":   "District",
		"City":       "City",
		"State":      "",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrStateMinLength)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
		"Street":     "Street ABC",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrAddressNumberMinLength)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
		"Street":     "Street ABC",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "zipcodee",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrZipcodeNotNumeric)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
		"Street":     "Street ABC",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrLatitudeInvalidFormat)

	err = json.Unmarshal([]byte(`
	{
		"id":         50,
		"Street":     "Street ABC",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  ""
	}
	`), &address)
	require.ErrorIs(err, uservo.ErrLongitudeInvalidFormat)

	err = json.Unmarshal([]byte(`
	{
		"id":         "",
		"Street":     "Street ABC",
		"District":   "District",
		"City":       "City",
		"State":      "State",
		"Complement": "Complement",
		"Number":     "11111",
		"Zipcode":    "23970000",
		"Latitude":   "-23.307577",
		"Longitude":  "-44.754146"
	}
	`), &address)
	require.Equal(http.StatusBadRequest, resperr.StatusCode(err))

}
