package hx

import (
	"fmt"
	"io"
	"net/http"
)

func ReadResponseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close() // Ensure the response body is closed after reading

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
