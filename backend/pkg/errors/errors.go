package errors

import (
	"fmt"
	"net/http"
)

type ErrorCode int

const (
	// 系统错误 (1xxxx)
	ErrSystem           ErrorCode = 10000
	ErrDatabase        ErrorCode = 10001
	ErrCache           ErrorCode = 10002
	ErrNetwork         ErrorCode = 10003
	
	// 业务错误 (2xxxx)
	ErrInvalidParam    ErrorCode = 20000
	ErrUnauthorized    ErrorCode = 20001
	ErrForbidden       ErrorCode = 20002
	ErrNotFound        ErrorCode = 20003
	
	// 实验错误 (3xxxx)
	ErrExperimentFailed ErrorCode = 30000
	ErrAgentNotFound    ErrorCode = 30001
	ErrAgentFailed      ErrorCode = 30002
)

type AppError struct {
	Code     ErrorCode `json:"code"`
	Message  string    `json:"message"`
	Details  string    `json:"details,omitempty"`
	Internal error     `json:"-"`
}

func (e *AppError) Error() string {
	if e.Internal != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Internal)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Internal
}

func NewAppError(code ErrorCode, message string, details string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

func (e *AppError) WithInternal(err error) *AppError {
	e.Internal = err
	return e
}

// HTTP状态码映射
func (e *AppError) HTTPStatusCode() int {
	switch e.Code {
	case ErrInvalidParam:
		return http.StatusBadRequest
	case ErrUnauthorized:
		return http.StatusUnauthorized
	case ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
} 