package day02

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	firstReq          int
	secondReq         int
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
	if err != nil || len(password) < rule.secondReq {
		return false, errors.New("Not valid password")
	}

	firstCondition := string(password[rule.firstReq-1]) == rule.requiredCharacter && string(password[rule.secondReq-1]) != rule.requiredCharacter
	secondCondition := string(password[rule.firstReq-1]) != rule.requiredCharacter && string(password[rule.secondReq-1]) == rule.requiredCharacter

	if firstCondition || secondCondition {
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
	firstReq, _ := strconv.Atoi(limits[0])
	secondReq, _ := strconv.Atoi(limits[1])
	requiredCharacter := ruleFields[1]

	return Rule{
		firstReq,
		secondReq,
		requiredCharacter,
	}, password, nil
}

func aoc02() string {
	return ""
}
