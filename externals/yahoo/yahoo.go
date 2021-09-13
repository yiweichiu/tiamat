package yahoo

import (
	"io"
	"net/http"
)

const baseUrl = "https://yfapi.net"
const apiKey = "pYhQXF3uNF5poQYvNMKcn7HPJq5SGcauaF7uci25"

func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-KEY", apiKey)
	return req, nil
}

func GetBaseUrl() string {
	return baseUrl
}

func GetApiKey() string {
	return apiKey
}
