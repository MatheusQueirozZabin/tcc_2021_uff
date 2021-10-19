package uservo

import (
	"errors"
	"ifoodish-store/pkg/resperr"
	"net/http"

	valid "github.com/asaskevich/govalidator"
)

var (
	ErrLatitudeInvalidFormat = errors.New(" latitude should have a valid format")
)

type Latitude string

func (s Latitude) Equals(other Latitude) bool {
	return s.String() == other.String()
}

func (s Latitude) String() string {
	return string(s)
}

func NewLatitude(value string) (Latitude, error) {
	if !valid.IsLatitude(value) {
		return "", resperr.WithCodeAndMessage(
			ErrLatitudeInvalidFormat,
			http.StatusBadRequest,
			"A Latitude é invalida, utilize o padrão 'X.XXXX' ou '-X.XXXX'",
		)
	}
	return Latitude(value), nil

}
