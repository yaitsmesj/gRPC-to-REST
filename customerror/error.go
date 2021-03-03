package customerror

import "fmt"

// StatusCodeError ...
type StatusCodeError struct {
	StatusCode int
	Body       []byte
}

func (s *StatusCodeError) Error() string {
	return fmt.Sprintf("Status Code: %v\nResponse Body: %v", s.StatusCode, s.Body)
}
