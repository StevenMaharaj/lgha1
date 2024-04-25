package tests

import (
	"net/http"
	"testing"
)

// make get request to localhost:8080
func getreq() *http.Response {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp

}

func TestLgha1(t *testing.T) {
	// Run server in a goroutine

	resp := getreq()
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

}
