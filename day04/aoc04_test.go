package day04

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var allFieldsPassport = "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\n" +
	"byr:1937 iyr:2017 cid:147 hgt:183cm"

var passportWithSevenFieldsAndCID = "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\n" +
	"hcl:#cfa07d byr:1929"

var passportWithSevenFieldsAndNoCID = "hcl:#ae17e1 iyr:2013\n" +
	"eyr:2024\n" +
	"ecl:brn pid:760753108 byr:1931\n" +
	"hgt:179cm"

var passportWithTwoMissingFields = "hcl:#cfa07d eyr:2025 pid:166559648\n" +
	"iyr:2011 ecl:brn hgt:59in"

func TestIsValidPassport(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
		err   error
	}{
		"Passport with all fields is valid": {
			input: allFieldsPassport, want: true, err: nil,
		},
		"Passport without one field is invalid and missing field isn't cid": {
			input: passportWithSevenFieldsAndCID, want: false, err: nil,
		},
		"Passport without one field is invalid if the missing field is cid": {
			input: passportWithSevenFieldsAndNoCID, want: true, err: nil,
		},
		"Passport with two missing fields is invalid": {
			input: passportWithTwoMissingFields, want: false, err: nil,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			isValid, err := IsValidPassport(testCase.input)

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

var exampleDataset = "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\n" +
	"byr:1937 iyr:2017 cid:147 hgt:183cm\n" +
	"\n" +
	"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\n" +
	"hcl:#cfa07d byr:1929\n" +
	"\n" +
	"hcl:#ae17e1 iyr:2013\n" +
	"eyr:2024\n" +
	"ecl:brn pid:760753108 byr:1931\n" +
	"hgt:179cm\n" +
	"\n" +
	"hcl:#cfa07d eyr:2025 pid:166559648\n" +
	"iyr:2011 ecl:brn hgt:59in"

func TestCountValidPassports(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
		err   error
	}{
		"Example dataset with 2 valid passports": {
			input: exampleDataset, want: 2, err: nil,
		},
		"Real dataset": {
			input: getRealDataset(), want: 237, err: nil,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			validPassports, err := CountValidPassports(testCase.input)

			if testCase.err != nil {
				assert.Error(t, err)
			}

			if testCase.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, testCase.want, validPassports)
		})
	}
}

func getRealDataset() string {
	const DatasetPath string = "data.txt"

	f, _ := os.Open(DatasetPath)

	reader := bufio.NewReader(f)
	// var dataset string
	var dataset []string
	for line := readLine(reader); line != nil; line = readLine(reader) {
		dataset = append(dataset, string(line))
	}

	return strings.Join(dataset, "\n")
}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}
