package errs

// CRUDエラー型
type CRUDError struct {
	message string
}

// getter
func (e *CRUDError) Error() string {
	return e.message
}

// コンストラクタ
func NewCRUDError(message string) *CRUDError {
	return &CRUDError{message: message}
}