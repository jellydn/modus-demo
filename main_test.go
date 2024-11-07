package main

import "testing"

func TestSayHello(t *testing.T) {
	// Create test cases using a table-driven testing approach
	tests := []struct {
		name     string  // test case name
		input    *string // input parameter
		expected string  // expected result
	}{
		{
			name:     "nil input returns Hello, World!",
			input:    nil,
			expected: "Hello, World!",
		},
		{
			name:     "with name returns personalized greeting",
			input:    stringPtr("Alice"),
			expected: "Hello, Alice!",
		},
		{
			name:     "empty string returns greeting with empty name",
			input:    stringPtr(""),
			expected: "Hello, !",
		},
	}

	// Run all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SayHello(tt.input)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

// Helper function to create string pointer
func stringPtr(s string) *string {
	return &s
}
