package errors_custom

import (
	"errors"
	"net/http"
)

var customErrorMap map[error]ErrorContent

var (
	// unknown error
	UnknownErrorContent = ErrorContent{
		HttpStatusCode:       500,
		ApplicationErrorCode: 90,
		CustomMessage:        "Some thing went wrong!",
	}
	// invalid credentials
	ErrInvalidCredentials = errors.New("invalid credentials")
	// empty token error
	ErrEmptyToken = errors.New("empty token")
	// invalid token
	ErrInvalidToken = errors.New("invalid token")
	// bad request
	ErrBadRequest = errors.New("bad request")
	// unable to add stakeholder
	ErrUnableToAddMobileAppUser = errors.New("unable to add mobile app user")
	// duplicate email
	ErrDuplicateUserEntry = errors.New("e-mail is already used")
)

func init() {
	customErrorMap = make(map[error]ErrorContent)
	// error description

	// invalid credentials
	customErrorMap[ErrInvalidCredentials] = ErrorContent{
		HttpStatusCode:       http.StatusUnauthorized,
		ApplicationErrorCode: 100,
		CustomMessage:        ErrInvalidCredentials.Error(),
	}
	// empty token error
	customErrorMap[ErrEmptyToken] = ErrorContent{
		HttpStatusCode:       http.StatusBadRequest,
		ApplicationErrorCode: 101,
		CustomMessage:        ErrEmptyToken.Error(),
	}
	// invalid token error
	customErrorMap[ErrInvalidToken] = ErrorContent{
		HttpStatusCode:       http.StatusUnauthorized,
		ApplicationErrorCode: 102,
		CustomMessage:        ErrInvalidToken.Error(),
	}
	// bad request
	customErrorMap[ErrBadRequest] = ErrorContent{
		HttpStatusCode:       http.StatusBadRequest,
		ApplicationErrorCode: 103,
		CustomMessage:        ErrBadRequest.Error(),
	}
	// unable to add  mobile app user
	customErrorMap[ErrUnableToAddMobileAppUser] = ErrorContent{
		HttpStatusCode:       http.StatusOK,
		ApplicationErrorCode: 201,
		CustomMessage:        ErrUnableToAddMobileAppUser.Error(),
	}
	// Duplicate email
	customErrorMap[ErrDuplicateUserEntry] = ErrorContent{
		HttpStatusCode:       http.StatusBadRequest,
		ApplicationErrorCode: 202,
		CustomMessage:        ErrDuplicateUserEntry.Error(),
	}
}
