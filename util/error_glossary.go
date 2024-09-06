package util

import "net/http"

type ErrorGlossary struct {
	ErrorCode      string
	ErrorMessage   string
	HTTPStatusCode int
}

var ERROR_GLOSSARY = map[string]ErrorGlossary{
	"ERR101": {
		ErrorCode:      "ERR101",
		ErrorMessage:   "Either username or email must be provided",
		HTTPStatusCode: http.StatusBadRequest,
	},
	"ERR102": {
		ErrorCode:      "ERR102",
		ErrorMessage:   "Unsupported Media Type",
		HTTPStatusCode: http.StatusUnsupportedMediaType,
	},
	"ERR103": {
		ErrorCode:      "ERR103",
		ErrorMessage:   "request body parse failed which expects a valid JSON structure",
		HTTPStatusCode: http.StatusBadRequest,
	},
	"ERR104": {
		ErrorCode:      "ERR104",
		ErrorMessage:   "Username or Email are not registerd.",
		HTTPStatusCode: http.StatusUnauthorized,
	},
	"ERR105": {
		ErrorCode:      "ERR105",
		ErrorMessage:   "Database Error",
		HTTPStatusCode: http.StatusInternalServerError,
	},
	"ERR106": {
		ErrorCode:      "ERR106",
		ErrorMessage:   "Password Encryption Error",
		HTTPStatusCode: http.StatusInternalServerError,
	},
	"ERR107": {
		ErrorCode:      "ERR107",
		ErrorMessage:   "Please enter correct password.",
		HTTPStatusCode: http.StatusUnauthorized,
	},
	"ERR108": {
		ErrorCode:      "ERR108",
		ErrorMessage:   "Failed to generate jwt token",
		HTTPStatusCode: http.StatusInternalServerError,
	},
}
