package interpreterutils

import "testing"

func TestInterpretInput(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Basic valid and invalid characters",
			input:    "+++strawberrycheesecake.[-]>>++++",
			expected: 13,
		},
		{
			name:     "Empty input",
			input:    "",
			expected: 0,
		},
		{
			name:     "Only valid characters",
			input:    "+++.[-]>>++++",
			expected: 13,
		},
		{
			name:     "Only invalid characters",
			input:    "strawberrycheesecake",
			expected: 0,
		},
		{
			name:     "Mixed valid characters",
			input:    ".,+-<>[]",
			expected: 8,
		},
		{
			name:     "Complex pattern",
			input:    "+++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.",
			expected: 103,
		},
		{
			name:     "Hello World!",
			input:    ">+++++++++[<++++++++>-]<.>+++++++[<++++>-]<+.+++++++..+++.>>>++++++++[<++++>-]\n<.>>>++++++++++[<+++++++++>-]<---.<<<<.+++.------.--------.>>+.>++++++++++.",
			expected: 153,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := interpretInput(tc.input)
			if len(output) != tc.expected {
				t.Errorf("Test '%s' failed: expected output length to be %d, got %d",
					tc.name, tc.expected, len(output))
			}
		})
	}

}
