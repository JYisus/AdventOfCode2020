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

func TestCountTreesInSlope(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
		err   error
	}{
		"The easiest map with just 1 tree in the slope": {
			input: easiestMapOfTrees, want: 1, err: nil,
		},
		"The easiest map without trees in the slope": {
			input: easiestMapOfTreesWithoutTrees, want: 0, err: nil,
		},
		"A real example map that have 7 trees in the slope": {
			input: exampleMap, want: 7, err: nil,
		},
		"Real dataset": {
			input: getRealDataset(), want: 191, err: nil,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			isThree, err := CountTreesInSlope(testCase.input)

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
