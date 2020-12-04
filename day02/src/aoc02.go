package day02

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	minRepetitions    int
	maxRepetitions    int
	requiredCharacter string
}

func CountValidPasswords(passwords []string) int {
	validPasswords := 0
	for _, password := range passwords {
		isValid, err := IsValidPassword(password)
		if err == nil && isValid {
			validPasswords++
		}
	}
	return validPasswords
}

func IsValidPassword(passwordLine string) (bool, error) {
	rule, password, err := ExtractRuleAndPassword(passwordLine)
	if err != nil {
		return false, errors.New("Not valid password")
	}

	repetitionsOfRequiredCharacter := strings.Count(password, rule.requiredCharacter)

	if repetitionsOfRequiredCharacter >= rule.minRepetitions && repetitionsOfRequiredCharacter <= rule.maxRepetitions {
		return true, nil
	}

	return false, nil

}

func ExtractRuleAndPassword(passwordLine string) (Rule, string, error) {
	matched, err := regexp.MatchString(`([0-9]+?)-([0-9]+?) (\w): (\w+?)`, passwordLine)
	if err != nil || !matched {
		return Rule{}, "", errors.New("Rule not valid")
	}

	passwordFields := strings.Split(passwordLine, ":")

	password := strings.Trim(passwordFields[1], " ")

	ruleFields := strings.Split(passwordFields[0], " ")
	limits := strings.Split(ruleFields[0], "-")
	minRepetitions, _ := strconv.Atoi(limits[0])
	maxRepetitions, _ := strconv.Atoi(limits[1])
	requiredCharacter := ruleFields[1]

	return Rule{
		minRepetitions,
		maxRepetitions,
		requiredCharacter,
	}, password, nil
}

func aoc02() string {
	return ""
}
