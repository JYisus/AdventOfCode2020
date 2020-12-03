package main

import (
	"log"

	"github.com/pkg/errors"
)

func AoC01(data []int) (int, error) {
	numbers, err := FindNumbers(data)
	if err != nil {
		return 0, err
	}

	return numbers[0] * numbers[1], err
}

func FindNumbers(data []int) ([]int, error) {

	if len(data) < 2 {
		return nil, errors.Errorf("Not valid input")
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data) && i != j; j++ {
			if data[i]+data[j] == 2020 {
				return []int{data[i], data[j]}, nil
			}
		}

	}

	return nil, errors.Errorf("Not valid input")
}

func main() {
	log.Println("Hello wdorld")
}
