package microservice

import "fmt"

func (c contextKey) String() string {
	return string(c)
}

func (c FilterType) Int32() int32 {
	return int32(c)
}

func (c msgError) String() string {
	return string(c)
}

func (c msgError) ToError() error {
	switch c {
	case msgInvalidToken:
		return ErrorInvalidToken
	case msgInvalidCode:
		return ErrorInvalidCode
	case msgInvalidRefreshToken:
		return ErrorInvalidRefreshToken
	case msgUnauthorizedAccess:
		return ErrUnauthorized
	case msgJWTExpired:
		return ErrJWTExpired
	case msgTokenEmpty:
		return ErrTokenEmpty
	default:
		return fmt.Errorf("Undefined the error")
	}
}
