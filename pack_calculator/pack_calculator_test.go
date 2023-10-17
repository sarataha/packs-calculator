package pack_calculator

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCalculatePacks(t *testing.T) {
	testCases := []struct {
		name     string
		items    int
		expected []Pack
	}{
		{
			name:  "1 item",
			items: 1,
			expected: []Pack{
				{Size: 250, Quantity: 1},
			},
		},
		{
			name:  "250 items",
			items: 250,
			expected: []Pack{
				{Size: 250, Quantity: 1},
			},
		},
		{
			name:  "251 items",
			items: 251,
			expected: []Pack{
				{Size: 500, Quantity: 1},
			},
		},
		{
			name:  "501 items",
			items: 501,
			expected: []Pack{
				{Size: 500, Quantity: 1},
				{Size: 250, Quantity: 1},
			},
		},
		{
			name:  "12001 items",
			items: 12001,
			expected: []Pack{
				{Size: 5000, Quantity: 2},
				{Size: 2000, Quantity: 1},
				{Size: 250, Quantity: 1},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			packs := CalculatePacks(tc.items)
			fmt.Println(packs)
			assert.Equal(t, tc.expected, packs)
		})
	}
}
