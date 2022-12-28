package internal

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientMovie(t *testing.T) {
	assert.Equal(t, 200, HttpServer())

	req, err := ClientMovie()
	assert.NoError(t, err, "ClientMovie returned an error")

	assert.Equal(t, "GET", req.Method, "Expected GET request")
	assert.Equal(t, "http://localhost:8080/movie", req.URL.String(), "Expected request to be sent to http://localhost:8080/movie")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err, "Error executing request")

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code of %d", http.StatusOK)
}
