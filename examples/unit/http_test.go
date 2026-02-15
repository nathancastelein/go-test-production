package unit

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHTTPCall(t *testing.T) {
	// Arrange
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/hello" {
			t.Fatalf("expected call, got %s on %s", r.Method, r.URL.Path)
		}

		fmt.Fprintln(w, "Hello world!")
	}))
	t.Cleanup(func() { testServer.Close() })

	// Act
	result, err := HelloHTTPCall(testServer.URL)

	// Assert
	if err != nil {
		t.Fatalf("got an error while making call: %s", err)
	}

	if result != "Hello world!\n" {
		t.Fatalf("expected body Hello world!, got %s", result)
	}
}

func HelloHTTPCall(url string) (string, error) {
	res, err := http.Get(url + "/hello")
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}

	return string(body), nil
}
