package helper

import (
	"bufio"
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

// Reads a file and returns each line in a string array, ignoring empty lines
func ReadLines(fh io.Reader, ignoreEmpty bool) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if ignoreEmpty && len(line) == 0 {
			continue
		}
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Reads a file of integers, one per line (ignoring empty), and returns the corresponding array
func ReadLinesAsInt(fh io.Reader) ([]int, error) {
	var lines []int
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		lines = append(lines, i)
	}
	return lines, scanner.Err()
}

// Reads a file of integers, one per line (ignoring empty), and returns the corresponding array
func ReadLinesAsIntArray(fh io.Reader) ([][]int, error) {
	var lines [][]int
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		iline := make([]int, 0)
		for _, r := range line {
			i, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}
			iline = append(iline, i)
		}
		lines = append(lines, iline)
	}
	return lines, scanner.Err()
}

// Reads a CSV file
func ReadCSV(fh io.Reader) ([][]string, error) {
	r := csv.NewReader(fh)
	return r.ReadAll()
}
