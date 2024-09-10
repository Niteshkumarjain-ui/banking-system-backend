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
	"ERR109": {
		ErrorCode:      "ERR109",
		ErrorMessage:   "Invalid Token",
		HTTPStatusCode: http.StatusUnauthorized,
	},
	"ERR110": {
		ErrorCode:      "ERR110",
		ErrorMessage:   "User is not authorized",
		HTTPStatusCode: http.StatusForbidden,
	},
	"ERR111": {
		ErrorCode:      "ERR111",
		ErrorMessage:   "User already has account.",
		HTTPStatusCode: http.StatusForbidden,
	},
	"ERR112": {
		ErrorCode:      "ERR112",
		ErrorMessage:   "Failed to retrieve account details.",
		HTTPStatusCode: http.StatusNotFound,
	},
	"ERR113": {
		ErrorCode:      "ERR113",
		ErrorMessage:   "Insufficent Balace.",
		HTTPStatusCode: http.StatusBadRequest,
	},
	"ERR114": {
		ErrorCode:      "ERR114",
		ErrorMessage:   "Transaction Not Found.",
		HTTPStatusCode: http.StatusNotFound,
	},
}
