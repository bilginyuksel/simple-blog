package main

// APIError ...
type APIError struct {
	code    int64
	message string
}

func (ap *APIError) Error() (int64, string) {
	return ap.code, ap.message
}
