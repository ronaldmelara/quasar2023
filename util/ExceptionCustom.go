package util

import(
	"fmt"
)

type Exception struct{
	StatusCode int
	Err error
}

func (m *Exception) Error() string{
	return fmt.Sprintf("status %d:  Message: %v", m.StatusCode, m.Err)
}