package interpreterutils

import "testing"

func TestExecuteBrainfuck(t *testing.T) {
	testCases := []struct {
		name           string
		input          []token
		expectedOutput string
	}{
		{name: "Hello World!",
			input: inputHelper(">+++++++++[<++++++++>-]<.>+++++++" +
				"[<++++>-]<+.+++++++..+++.>>>++++++++[<++++>-]" +
				"<.>>>++++++++++[<+++++++++>-]" +
				"<---.<<<<.+++.------.--------.>>+.>++++++++++."),
			expectedOutput: "Hello World!",
		},
	}
	for _, tc := range testCases {
		output, _ := executeBrainfuck(tc.input)
		if output != tc.expectedOutput {
			t.Errorf("Test '%s' failed: expected output to be %s, got %s",
				tc.name, tc.expectedOutput, output)
		}

	}
}

func inputHelper(input string) []token {
	tokens := make([]token, 0, len(input))

	for _, char := range input {
		tokens = append(tokens, token{symbol: string(char)})
	}

	return tokens
}
