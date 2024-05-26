package util

import (
	"errors"

	"github.com/muharanii/auth-rest-api-go/domain"
)

func GetHttpStatus(err error) int {
	switch{
	case errors.Is(err, domain.ErrAuthFailed):
		return 401
	default:
		return 500

	}
}