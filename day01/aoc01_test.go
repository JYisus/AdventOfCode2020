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
		"given no numbers, return an error":                  {input: nil, want: nil, err: errors.New("Not valid input")},
		"given just one value returns an error":              {input: []int{1}, want: nil, err: errors.New("Not valid input")},
		"given only two valid values return an error":        {input: []int{2020, 0}, want: nil, err: errors.New("Not valid input")},
		"given three valid values returns all of them":       {input: []int{2017, 1, 2}, want: []int{2017, 1, 2}, err: nil},
		"given four values returns the three that sums 2020": {input: []int{2010, 1, 3, 7}, want: []int{2010, 3, 7}, err: nil},
		"given four values returns an error " +
			"if there isn't a threesome that sums 2020": {input: []int{2010, 100, 3, 12}, want: nil, err: errors.New("Not valid input")},
		"given a biggest amout of numbers returns the values that sums 2020": {input: exampleTestValues, want: []int{979, 366, 675}, err: nil},
		// "given the real dataset returns the pair of values that sums 2020":   {input: getRealDataset(), want: []int{529, 1491}, err: nil},
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
		"given no numbers, return an error":                  {input: nil, want: 0, err: errors.New("Not valid input")},
		"given just one value returns an error":              {input: []int{1}, want: 0, err: errors.New("Not valid input")},
		"given two valid values returns error":               {input: []int{2020, 0}, want: 0, err: errors.New("Not valid input")},
		"given three invalid values returns it mults":        {input: []int{2017, 1, 2}, want: 4034, err: nil},
		"given four values returns the three that sums 2020": {input: []int{2010, 1, 3, 7}, want: 42210, err: nil},
		"given four values returns an error " +
			"if there isn't a threesome that sums 2020": {input: []int{2010, 100, 3, 12}, want: 0, err: errors.New("Not valid input")},
		"given a biggest amout of numbers returns the values that sums 2020": {input: exampleTestValues, want: 241861950, err: nil},
		"given the real dataset returns the final result":                    {input: getRealDataset(), want: 178724430, err: nil},
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
