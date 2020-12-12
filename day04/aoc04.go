package day04

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func CountValidPassports(dataset string) (int, error) {
	passports := strings.Split(dataset, "\n\n")

	validPassports := 0

	for _, passport := range passports {
		isValid, _ := IsValidPassport(passport)
		if isValid {
			fmt.Println("Valid:")
			fmt.Println(passport)
			validPassports++
		}

	}
	return validPassports, nil
}

func IsValidPassport(passport string) (bool, error) {
	fieldsLines := strings.Split(passport, "\n")

	fields := []string{}
	for _, line := range fieldsLines {
		fieldsInLine := strings.Split(line, " ")
		fields = append(fields, fieldsInLine...)
	}

	if len(fields) < 7 {
		return false, nil
	}

	haveCid := false
	validFields := 0
	for _, field := range fields {
		splited := strings.Split(field, ":")
		key := splited[0]
		value := splited[1]

		switch key {
		case "cid":
			// haveCid = true
			validFields++
		case "byr":

			valueNumber, err := strconv.Atoi(value)
			if err != nil {
				return false, nil
			}
			if valueNumber < 1920 || valueNumber > 2002 {
				return false, nil
			}
			validFields++

		case "iyr":

			valueNumber, err := strconv.Atoi(value)
			if err != nil {
				return false, nil
			}
			if valueNumber < 2010 || valueNumber > 2020 {
				return false, nil
			}
			validFields++

		case "eyr":

			valueNumber, err := strconv.Atoi(value)
			if err != nil {
				return false, nil
			}
			if valueNumber < 2020 || valueNumber > 2030 {
				return false, nil
			}
			validFields++

		case "hgt":

			unit := value[len(value)-2 : len(value)]
			hgt, err := strconv.Atoi(value[:len(value)-2])

			if err != nil {
				return false, nil
			}

			if unit == "cm" {
				if hgt < 150 || hgt > 193 {
					return false, nil
				}
			} else {
				if unit == "in" {
					if hgt < 59 || hgt > 76 {
						return false, nil
					}
				} else {
					return false, nil
				}
			}
			validFields++

		case "hcl":
			matched, _ := regexp.MatchString(`#([0-9a-f]{6,6})`, value)
			if !matched {
				return false, nil
			}
			validFields++

		case "ecl":
			matched, _ := regexp.MatchString(`amb|blu|brn|gry|grn|hzl|oth`, value)
			if !matched {
				return false, nil
			}
			validFields++

		case "pid":
			matched, _ := regexp.MatchString(`([0-9]{9,9})`, value)
			if !matched {
				return false, nil
			}
			validFields++

		}

	}

	if validFields == 8 {
		return true, nil
	}

	if validFields == 7 && haveCid == false {
		return true, nil
	}

	return false, nil
}
