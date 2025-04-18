package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts the API key from the request headers and returns it.
// ErrMissingAPIKey is returned when the API key is not found in the request headers.
// Header example:
// x-api-key: ApiKey <API_KEY>
func GetAPIKey(headers http.Header) (string, error) {
	apiKey := headers.Get("x-api-key")
	if apiKey == "" {
		return "", errors.New("missing api key")
	}

	apiKeyStrings := strings.Split(apiKey, " ")
	if len(apiKeyStrings) != 2 {
		return "", errors.New("malformed api key")
	}

	if apiKeyStrings[0] != "ApiKey" {
		return "", errors.New("malformed api key")
	}

	return apiKeyStrings[1], nil
}
