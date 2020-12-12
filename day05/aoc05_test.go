package day05

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var validPassport = "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\n" +
	"hcl:#623a2f"

func TestCheckRow(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
		err   error
	}{
		"Row 44 specification": {
			input: "FBFBBFFRLR", want: 44, err: nil,
		},
		// "Row 70 spec": {
		// 	input: "BFFFBBFRRR", want: 70, err: nil,
		// },
		// "Row 14 spec": {
		// 	input: "FFFBBBFRRR", want: 14, err: nil,
		// },
		// "Row 102 spec": {
		// 	input: "BBFFBBFRLL", want: 102, err: nil,
		// },
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			rowNumber, err := CheckRow(testCase.input)

			if testCase.err != nil {
				assert.Error(t, err)
			}

			if testCase.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, testCase.want, rowNumber)
		})
	}
}

func TestCheckColumn(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
		err   error
	}{
		"Column 44 specification": {
			input: "FBFBBFFRLR", want: 5, err: nil,
		},
		"Column 70 spec": {
			input: "BFFFBBFRRR", want: 7, err: nil,
		},
		"Column 14 spec": {
			input: "FFFBBBFRRR", want: 7, err: nil,
		},
		"Column 102 spec": {
			input: "BBFFBBFRLL", want: 4, err: nil,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			ColumnNumber, err := CheckColumn(testCase.input)

			if testCase.err != nil {
				assert.Error(t, err)
			}

			if testCase.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, testCase.want, ColumnNumber)
		})
	}
}

func TestSeatID(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
		err   error
	}{
		"Column 44 specification": {
			input: "FBFBBFFRLR", want: 357, err: nil,
		},
		"Column 70 spec": {
			input: "BFFFBBFRRR", want: 567, err: nil,
		},
		"Column 14 spec": {
			input: "FFFBBBFRRR", want: 119, err: nil,
		},
		"Column 102 spec": {
			input: "BBFFBBFRLL", want: 820, err: nil,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			seatID, err := SeatID(testCase.input)

			if testCase.err != nil {
				assert.Error(t, err)
			}

			if testCase.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, testCase.want, seatID)
		})
	}
}

func TestAoC05(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  int
		err   error
	}{
		"Real dataset": {
			input: getRealDataset(), want: 935, err: nil,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			maxID, err := AoC05(testCase.input)

			if testCase.err != nil {
				assert.Error(t, err)
			}

			if testCase.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, testCase.want, maxID)
		})
	}
}

func getRealDataset() []string {
	const DatasetPath string = "data.txt"

	f, _ := os.Open(DatasetPath)

	reader := bufio.NewReader(f)
	// var dataset string
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
