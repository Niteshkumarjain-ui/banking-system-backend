package domain

type HTTPError struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
