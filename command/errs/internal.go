package errs

// internalエラー型
type InternalError struct {
	message string
}

// getter
func (e *InternalError) Error() string {
	return e.message
}

// コンストラクタ
func NewInternalError(message string) *InternalError {
	return &InternalError{message: message}
}