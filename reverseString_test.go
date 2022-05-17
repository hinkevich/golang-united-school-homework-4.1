package reverse_string

import "testing"

type testData struct {
	inputString string
	want        string
}

func TestReveseString(t *testing.T) {

	tests := []testData{
		{inputString: "abcdef", want: "fedcba"},
		{inputString: "☹abcdef", want: "fedcba☹"},
		{inputString: "abc😎def", want: "fed😎cba"},
	}
	for i, test := range tests {
		got := ReverseString(test.inputString)
		if got != test.want {
			t.Errorf("test № %d is failed ", i)
		}
	}
}
