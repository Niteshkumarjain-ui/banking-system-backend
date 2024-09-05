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
		ErrorMessage:   "Upstream service unavailable",
		HTTPStatusCode: http.StatusServiceUnavailable,
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
		ErrorMessage:   "Trailing Space Error",
		HTTPStatusCode: http.StatusBadRequest,
	},
	"ERR105": {
		ErrorCode:      "ERR105",
		ErrorMessage:   "Database Error",
		HTTPStatusCode: http.StatusInternalServerError,
	},
}
