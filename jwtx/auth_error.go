package jwtx

import (
	"errors"
)

var (
	ErrEmptyAuthHeader         = errors.New("auth header is empty")
	ErrEmptyQueryToken         = errors.New("query token is empty")
	ErrEmptyCookieToken        = errors.New("cookie token is empty")
	ErrEmptyParamToken         = errors.New("parameter token is empty")
	ErrInvalidAuthHeader       = errors.New("auth header is invalid")
	ErrInvalidSigningAlgorithm = errors.New("invalid signing algorithm")
	ErrExpiredToken            = errors.New("token is expired")
)
