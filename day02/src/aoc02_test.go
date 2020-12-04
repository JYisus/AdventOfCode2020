package day02

import (
	"bufio"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getRealDataset() []string {
	const DatasetPath string = "data.txt"

	f, _ := os.Open(DatasetPath)

	reader := bufio.NewReader(f)
	var dataset []string
	for line := readLine(reader); line != nil; line = readLine(reader) {
		dataset = append(dataset, string(line))
	}
	return dataset
}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}

func TestExtractRuleAndPassword(t *testing.T) {
	tests := map[string]struct {
		input        string
		wantRule     Rule
		wantPassword string
		err          error
	}{
		"a password that requires 1 to 2 'a' and have an 'a' return the rule and the password": {
			input:        "1-2 a: abcde",
			wantRule:     Rule{minRepetitions: 1, maxRepetitions: 2, requiredCharacter: "a"},
			wantPassword: "abcde",
			err:          nil},
		"a password that requires 1 to 3 'a' and have an 'a' return the rule and the password": {
			input:        "2-3 b: abcde",
			wantRule:     Rule{minRepetitions: 2, maxRepetitions: 3, requiredCharacter: "b"},
			wantPassword: "abcde",
			err:          nil},
		"a password with no requirements returns an error": {
			input:        "abcde",
			wantRule:     Rule{},
			wantPassword: "",
			err:          errors.New("Rule not valid"),
		},
		"a bad formed password returns an error": {
			input:        "a-2 d: abcde",
			wantRule:     Rule{},
			wantPassword: "",
			err:          errors.New("Rule not valid")},
		"another bad formed password returns an error": {
			input:        "a-2: abcde",
			wantRule:     Rule{},
			wantPassword: "",
			err:          errors.New("Rule not valid")},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			rule, password, err := ExtractRuleAndPassword(testCase.input)

			if testCase.err != nil {
				assert.Error(t, err)
			}

			if testCase.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, testCase.wantRule, rule)
			assert.Equal(t, testCase.wantPassword, password)
		})
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
		err   error
	}{
		"a password that requires 1 to 2 'a' and have an 'a' is valid": {
			input: "1-2 a: abcde",
			want:  true,
			err:   nil,
		},
		"a password that requires 1 to 2 'a' and have more 'a' is not valid": {
			input: "1-2 a: aaabcde",
			want:  false,
			err:   nil,
		},
		"a password that requires 1 to 2 'a' and have no 'a' is not valid": {
			input: "1-2 a: bcde",
			want:  false,
			err:   nil,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			isValid, err := IsValidPassword(testCase.input)

			if testCase.err != nil {
				assert.Error(t, err)
			}

			if testCase.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, testCase.want, isValid)
		})
	}
}

func TestCountValidPasswords(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  int
	}{
		"set with 1 valid password": {
			input: []string{"1-2 a: abcde"},
			want:  1,
		},
		"set with 1 not valid password": {
			input: []string{"1-2 a: aaabcde"},
			want:  0,
		},
		"set with 1 valid password and 1 not valid password": {
			input: []string{"1-2 a: aaabcde", "1-2 a: abcde"},
			want:  1,
		},
		"a bigger set of passwords": {
			input: []string{
				"1-3 a: abcde",
				"1-3 b: cdefg",
				"2-9 c: ccccccccc",
			},
			want: 2,
		},
		"real dataset": {
			input: getRealDataset(),
			want:  458,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			isValid := CountValidPasswords(testCase.input)

			assert.Equal(t, testCase.want, isValid)
		})
	}
}
