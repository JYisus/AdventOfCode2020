package day04

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var validPassport = "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\n" +
	"hcl:#623a2f"

var invalidPassport1 = "iyr:2019\n" +
	"hcl:#602927 eyr:1967 hgt:170cm\n" +
	"ecl:grn pid:012533040 byr:1946"

var validPassport2 = "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"

var invalidPassport = "eyr:1972 cid:100\n" +
	"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926"

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
			input: validPassport, want: true, err: nil,
		},
		"Passport without one field is invalid and missing field isn't cid": {
			input: invalidPassport, want: false, err: nil,
		},
		"another invalid passport": {
			input: invalidPassport1, want: false, err: nil,
		},
		// "Passport without one field is invalid if the missing field is cid": {
		// 	input: passportWithSevenFieldsAndNoCID, want: true, err: nil,
		// },
		// "Passport with two missing fields is invalid": {
		// 	input: passportWithTwoMissingFields, want: false, err: nil,
		// },
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

var someInvalidPassports = "eyr:1972 cid:100\n" +
	"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926\n" +
	"\n" +
	"iyr:2019\n" +
	"hcl:#602927 eyr:1967 hgt:170cm\n" +
	"ecl:grn pid:012533040 byr:1946\n" +
	"\n" +
	"hcl:dab227 iyr:2012\n" +
	"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277\n" +
	"\n" +
	"hgt:59cm ecl:zzz\n" +
	"eyr:2038 hcl:74454a iyr:2023\n" +
	"pid:3556412378 byr:2007"

var someValidPassports = "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\n" +
	"hcl:#623a2f\n" +
	"\n" +
	"eyr:2029 ecl:blu cid:129 byr:1989\n" +
	"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm\n" +
	"\n" +
	"hcl:#888785\n" +
	"hgt:164cm byr:2001 iyr:2015 cid:88\n" +
	"pid:545766238 ecl:hzl\n" +
	"eyr:2022\n" +
	"\n" +
	"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"

func TestCountValidPassports(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
		err   error
	}{
		// "Example dataset with 2 valid passports": {
		// 	input: exampleDataset, want: 2, err: nil,
		// },
		// "4 invalid passports": {
		// 	input: someInvalidPassports, want: 0, err: nil,
		// },
		// "4 valid passports": {
		// 	input: someValidPassports, want: 4, err: nil,
		// },
		"Real dataset": {
			input: getRealDataset(), want: 184, err: nil,
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
