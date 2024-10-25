package main

import "testing"

func TestMainAlgorithm(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		shouildFail bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true}, // Некорректная строка, должен быть failure
		{"", "", false},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			output, err := UnpackString(tc.input)
			if tc.shouildFail {
				if err == nil {
					t.Errorf("Expected failure for input %q, but got %v with output %q", tc.input, err, output)
				}
				return
			}
			if err != nil {
				t.Errorf("Expected no error for input %q, but got: %q", tc.input, err)
				return
			}
			if output != tc.expected {
				t.Errorf("For input %q, expected %q, but got %q", tc.input, tc.expected, output)
			}
		})
	}
}
