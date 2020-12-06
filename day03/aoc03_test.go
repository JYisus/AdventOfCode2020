package day03

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var easiestMapOfTrees string = "....\n" +
	"...#"

var easiestMapOfTreesWithoutTrees string = "....\n" +
	"...."

var exampleMap string = "..##.......\n" +
	"#...#...#..\n" +
	".#....#..#.\n" +
	"..#.#...#.#\n" +
	".#...##..#.\n" +
	"..#.##.....\n" +
	".#.#.#....#\n" +
	".#........#\n" +
	"#.##...#...\n" +
	"#...##....#\n" +
	".#..#...#.#"

func TestCountTreesInSlope3_1(t *testing.T) {
	tests := map[string]struct {
		inputMap  string
		inputStep Step
		want      int
		err       error
	}{
		"The easiest map with just 1 tree in the slope": {
			inputMap: easiestMapOfTrees, inputStep: Step{right: 3, down: 1}, want: 1, err: nil,
		},
		"The easiest map without trees in the slope": {
			inputMap: easiestMapOfTreesWithoutTrees, inputStep: Step{right: 3, down: 1}, want: 0, err: nil,
		},
		"A real example map that have 7 trees in the slope": {
			inputMap: exampleMap, inputStep: Step{right: 3, down: 1}, want: 7, err: nil,
		},
		"Real dataset": {
			inputMap: getRealDataset(), inputStep: Step{right: 3, down: 1}, want: 191, err: nil,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			isThree, err := CountTreesInSlope(testCase.inputMap, testCase.inputStep)

			if testCase.err != nil {
				assert.Error(t, err)
			}

			if testCase.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, testCase.want, isThree)
		})
	}
}

func TestCountTreesWithDifferentSlopes(t *testing.T) {
	tests := map[string]struct {
		inputMap  string
		inputStep Step
		want      int
		err       error
	}{
		"The easiest map with 1x1 slope": {
			inputMap: easiestMapOfTrees, inputStep: Step{right: 1, down: 1}, want: 0, err: nil,
		},
		"The easiest map with 5x1": {
			inputMap: easiestMapOfTreesWithoutTrees, inputStep: Step{right: 5, down: 1}, want: 0, err: nil,
		},
		"The easiest map with 7x1": {
			inputMap: easiestMapOfTreesWithoutTrees, inputStep: Step{right: 7, down: 1}, want: 0, err: nil,
		},
		"The easiest map with 1x2": {
			inputMap: easiestMapOfTreesWithoutTrees, inputStep: Step{right: 1, down: 2}, want: 0, err: nil,
		},
		"A real example map with 1x1 slope": {
			inputMap: exampleMap, inputStep: Step{right: 1, down: 1}, want: 2, err: nil,
		},
		"A real example map with 3x1 slope": {
			inputMap: exampleMap, inputStep: Step{right: 3, down: 1}, want: 7, err: nil,
		},
		"A real example map with 5x1 slope": {
			inputMap: exampleMap, inputStep: Step{right: 5, down: 1}, want: 3, err: nil,
		},
		"A real example map with 7x1 slope": {
			inputMap: exampleMap, inputStep: Step{right: 7, down: 1}, want: 4, err: nil,
		},
		"A real example map with 1x2 slope": {
			inputMap: exampleMap, inputStep: Step{right: 1, down: 2}, want: 2, err: nil,
		},
		"Real dataset": {
			inputMap: getRealDataset(), inputStep: Step{right: 3, down: 1}, want: 191, err: nil,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			isThree, err := CountTreesInSlope(testCase.inputMap, testCase.inputStep)

			if testCase.err != nil {
				assert.Error(t, err)
			}

			if testCase.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, testCase.want, isThree)
		})
	}
}

var slopes []Step = []Step{
	Step{right: 1, down: 1},
	Step{right: 3, down: 1},
	Step{right: 5, down: 1},
	Step{right: 7, down: 1},
	Step{right: 1, down: 2},
}

func TestAoC03(t *testing.T) {
	tests := map[string]struct {
		inputMap    string
		inputSlopes []Step
		want        int
		err         error
	}{
		"The easiest map with slopes": {
			inputMap: easiestMapOfTrees, inputSlopes: slopes, want: 0, err: nil,
		},
		"A real example map with slopes": {
			inputMap: exampleMap, inputSlopes: slopes, want: 336, err: nil,
		},
		"Real dataset": {
			inputMap: getRealDataset(), inputSlopes: slopes, want: 1478615040, err: nil,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			isThree, err := AoC03(testCase.inputMap, testCase.inputSlopes)

			if testCase.err != nil {
				assert.Error(t, err)
			}

			if testCase.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, testCase.want, isThree)
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

	fmt.Println("dataset")
	return strings.Join(dataset, "\n")
}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}
