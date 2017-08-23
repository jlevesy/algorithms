package strings

import (
	"testing"
)

func TestHasUniqueChars(t *testing.T) {
	cases := []struct {
		Label       string
		Input       string
		Expectation bool
	}{
		{"WithUniqueChars", "abcdef", true},
		{"WithNonUniqueChars", "aabcdef", false},
		{"WithEmptyString", "", true},
		{"WithWhiteSpaces", "abcde ", true}, {"WithNonASCIIChars", "ABCèèèèèéééé", false}}

	for _, test := range cases {
		t.Run(test.Label, func(t *testing.T) {
			if test.Expectation != HasUniqueChars(test.Input) {
				t.Error("Algo 1: Unexpected result for input:", test.Input)
			}
		})
	}

	for _, test := range cases {
		t.Run(test.Label, func(t *testing.T) {
			if test.Expectation != HasUniqueChars2(test.Input) {
				t.Error("Algo 2 : Unexpected result for input:", test.Input)
			}
		})
	}
}

func TestIsPermutationOf(t *testing.T) {
	cases := []struct {
		Label       string
		Input1      string
		Input2      string
		Expectation bool
	}{
		{"WithPermutation", "abcdef", "fedcba", true},
		{"WithNoPermutation", "abcdef", "fegdcba", false},
		{"WithDifferentLength", "abcdef", "abcde", false},
		{"WithNonASCIIChars", "éàçèèè", "èèéçàè", true},
		{"WithEmptyStrings", "", "", true},
	}

	for _, test := range cases {
		t.Run(test.Label, func(t *testing.T) {
			if test.Expectation != IsPermutationOf(test.Input1, test.Input2) {
				t.Error("Unexpected results for:", test.Input1, test.Input2)
			}
		})
	}
}

func TestReplaceSpaces(t *testing.T) {
	cases := []struct {
		Label       string
		Input       string
		InputLength int
		ShouldPanic bool
		Expectation string
	}{
		{"WithNominalCase", "John Doe is Awesome      ", 19, false, "John%20Doe%20is%20Awesome"},
		{"WithInvalidLength", "John Doe is Awesome      ", 20, true, ""},
	}

	for _, test := range cases {
		t.Run(test.Label, func(t *testing.T) {
			if test.ShouldPanic {
				defer func() {
					if err := recover(); err == nil {
						t.Error("Code was expected to panic")
					}
				}()
			}

			if test.Expectation != string(ReplaceSpaces([]rune(test.Input), test.InputLength)) {
				t.Error("Unexpected result for input", test.Input, test.Expectation)
			}
		})
	}
}

func TestIsEquilibrated(t *testing.T) {
	cases := []struct {
		Label       string
		Input       string
		Chars       map[[2]rune]uint
		Expectation bool
	}{
		{"WithEquilibratedString", "((([[])))", map[[2]rune]uint{{'(', ')'}: 0}, true},
		{"WithNotEquilibratedString", "((([[])))", map[[2]rune]uint{{'(', ')'}: 0, {'[', ']'}: 0}, false},
		{"WithSmiley", "(((:[[])))", map[[2]rune]uint{{'(', ')'}: 0, {'[', ']'}: 0}, true},
		{"WithEmptyString", "", map[[2]rune]uint{{'(', ')'}: 0, {'[', ']'}: 0}, true},
		{"WithNonASCIIChars", "éèééèèéè", map[[2]rune]uint{{'é', 'è'}: 0}, true},
	}

	for _, test := range cases {
		t.Run(test.Label, func(t *testing.T) {
			if test.Expectation != IsEquilibrated(test.Input, test.Chars) {
				t.Error("Unexpected result for input", test.Input, test.Chars)
			}
		})
	}
}

func TestCompressString(t *testing.T) {
	compresedLengthCases := []struct {
		Label       string
		Input       string
		Expectation int
	}{
		{"WithLessThan9Characters", "aaaabbbbcccc", 6},
		{"WithMoreThan9Characters", "aaaaaaaaaa", 3},
	}

	for _, test := range compresedLengthCases {
		t.Run(test.Label, func(t *testing.T) {
			t.Log(compressedLength([]rune(test.Input)))
			if test.Expectation != compressedLength([]rune(test.Input)) {
				t.Error("Invalid result for input:", test.Input)
			}
		})
	}

	cases := []struct {
		Label       string
		Input       string
		Expectation string
	}{
		{"WithEfficientCompression", "aaaaaabbbbbbccc", "a6b6c3"},
		{"WithInefficientCompression", "ab", "ab"},
		{"WithEmptyString", "", ""},
		{"WitNonASCIIChars", "ééééé3333ççç@@@@àààà", "é534ç3@4à4"},
	}

	for _, test := range cases {
		t.Run(test.Label, func(t *testing.T) {
			t.Log(string(CompressString([]rune(test.Input))))
			if test.Expectation != string(CompressString([]rune(test.Input))) {
				t.Error("Unexpected result for input", test.Expectation)
			}
		})
	}
}

func BenchmarkCompressSting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompressString([]rune("aaaaaaabbbbbbbbbccccccccfffffffeeeeeee"))
	}
}
