package day04

import (
	"strings"
)

func CountValidPassports(dataset string) (int, error) {
	passports := strings.Split(dataset, "\n\n")

	validPassports := 0
	for _, passport := range passports {
		isValid, _ := IsValidPassport(passport)
		if isValid {
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

	if len(fields) == 8 {
		return true, nil
	}

	if len(fields) < 7 {
		return false, nil
	}

	haveCid := false
	for _, field := range fields {
		key := strings.Split(field, ":")[0]
		if key == "cid" {
			haveCid = true
		}
	}

	if haveCid == false {
		return true, nil
	}

	return false, nil
}
