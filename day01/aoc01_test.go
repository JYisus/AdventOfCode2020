package day01

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleTestValues = []int{
	1721,
	979,
	366,
	299,
	675,
	1456,
}

func getRealDataset() []int {
	const DatasetPath string = "data.txt"

	f, _ := os.Open(DatasetPath)

	reader := bufio.NewReader(f)
	var dataset []int
	for line := readLine(reader); line != nil; line = readLine(reader) {
		number, _ := strconv.Atoi(string(line))
		dataset = append(dataset, number)
	}
	return dataset
}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}

func TestFindNumbers(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
		err   error
	}{
		"given no numbers, return an error":                 {input: nil, want: nil, err: errors.New("Not valid input")},
		"given just one value returns an error":             {input: []int{1}, want: nil, err: errors.New("Not valid input")},
		"given two valid values return both of them":        {input: []int{2020, 0}, want: []int{2020, 0}, err: nil},
		"given two invalid values returns an error":         {input: []int{2022, 1}, want: nil, err: errors.New("Not valid input")},
		"given three values returns the two that sums 2020": {input: []int{2010, 3, 10}, want: []int{2010, 10}, err: nil},
		"given three values returns an error " +
			"if there isn't a pair that sums 2020": {input: []int{2010, 3, 12}, want: nil, err: errors.New("Not valid input")},
		"given a biggest amout of numbers returns the values that sums 2020": {input: exampleTestValues, want: []int{1721, 299}, err: nil},
		"given the real dataset returns the pair of values that sums 2020":   {input: getRealDataset(), want: []int{529, 1491}, err: nil},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := FindNumbers(testCase.input)

			if testCase.err != nil {
				assert.Error(t, err)
			}

			if testCase.err == nil {
				assert.Nil(t, err)
			}

			assert.ElementsMatch(t, testCase.want, result)
		})
	}
}

func TestAoC01(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
		err   error
	}{
		"given no numbers, return an error":                 {input: nil, want: 0, err: errors.New("Not valid input")},
		"given just one value returns an error":             {input: []int{1}, want: 0, err: errors.New("Not valid input")},
		"given two valid values returns their product":      {input: []int{2020, 0}, want: 0, err: nil},
		"given two invalid values returns an error":         {input: []int{2022, 1}, want: 0, err: errors.New("Not valid input")},
		"given three values returns the two that sums 2020": {input: []int{2010, 3, 10}, want: 20100, err: nil},
		"given three values returns an error " +
			"if there isn't a pair that sums 2020": {input: []int{2010, 3, 12}, want: 0, err: errors.New("Not valid input")},
		"given a biggest amout of numbers returns the values that sums 2020": {input: exampleTestValues, want: 514579, err: nil},
		"given the real dataset returns the final result":                    {input: getRealDataset(), want: 788739, err: nil},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := AoC01(testCase.input)

			if testCase.err != nil {
				assert.Error(t, err)
			}

			if testCase.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, testCase.want, result)
		})
	}
}
