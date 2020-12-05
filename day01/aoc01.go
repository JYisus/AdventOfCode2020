package day01

import (
	"github.com/pkg/errors"
)

func AoC01(data []int) (int, error) {
	numbers, err := FindNumbers(data)
	if err != nil {
		return 0, err
	}

	return numbers[0] * numbers[1] * numbers[2], err
}

func FindNumbers(data []int) ([]int, error) {

	if len(data) < 2 {
		return nil, errors.Errorf("Not valid input")
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data) && i != j; j++ {
			sum := data[i] + data[j]
			if sum > 2020 {
				continue
			}
			for k := 0; k < len(data) && i != k && j != k; k++ {
				if sum+data[k] == 2020 {
					return []int{data[i], data[j], data[k]}, nil
				}
			}
		}

	}

	return nil, errors.Errorf("Not valid input")
}
