package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculatePacksHandler(t *testing.T) {
	testCases := []struct {
		name     string
		items    string
		expected string
		status   int
	}{
		// Valid input
		{
			name:     "ValidInput",
			items:    "10",
			expected: `{"Items":10,"Packs":[{"Quantity":1,"Size":250}]}`,
			status:   http.StatusOK,
		},
		// Invalid input: items not a valid integer
		{
			name:     "InvalidInput",
			items:    "abc",
			expected: `Invalid input: items must be a valid integer`,
			status:   http.StatusBadRequest,
		},
		// Zero items
		{
			name:     "ZeroItems",
			items:    "0",
			expected: `{"Items":0,"Packs":[]}`,
			status:   http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/?items="+tc.items, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			PacksHandler(rr, req)

			if rr.Code != tc.status {
				t.Errorf("Expected status code %d, but got %d", tc.status, rr.Code)
			}

			if strings.TrimSpace(rr.Body.String()) != tc.expected {
				t.Errorf("Expected response:\n%s\nBut got:\n%s", tc.expected, rr.Body.String())
			}
		})
	}
}
