package storage

type ErrStorage struct {
	Code    int
	Message string
}

func CreateNew(code int, msg string) *ErrStorage {
	return &ErrStorage{
		Code:    code,
		Message: msg,
	}
}

func FromError(err error) *ErrStorage {
	return &ErrStorage{
		Code:    502,
		Message: err.Error(),
	}
}
