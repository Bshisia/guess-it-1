package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadFile(filename string) ([]float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	var result []float64
	count := 0
	for _, line := range lines {
		count++
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		val, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Printf("error: non-integer value %v incorporated at line %d\n", err, count)
			continue
		}
		result = append(result, val)
	}
	return result, nil
}

func Average(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

func Median(data []float64) float64 {
	sort.Float64s(data)
	len := len(data)
	if len%2 == 0 {
		return (data[(len-1)/2] + data[len/2]) / 2.0
	}
	return data[len/2]
}

func Variance(data []float64) float64 {
	average := Average(data)
	var sum float64
	for _, value := range data {
		difference := value - average
		sum += difference * difference
	}
	return sum / float64(len(data))
}

func StandardDeviation(data []float64) float64 {
	variance := Variance(data)
	return math.Sqrt(variance)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run [program name] [filename]")
		os.Exit(0)
	}
	filename := os.Args[1]
	if !strings.HasSuffix(filename, ".txt") {
		fmt.Println("Filename should be a txt file")
		os.Exit(1)
	}

	data, err := ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}
	if len(data) == 0 {
		fmt.Println("Empty data file")
		os.Exit(1)
	}

	// Create a scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter numbers (one per line):")

	for scanner.Scan() {
		line := scanner.Text()
		val, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Printf("Invalid input: %v\n", err)
			continue
		}
		data = append(data, val)

		average := Average(data)
		stdDev := StandardDeviation(data)

		// Predict the range for the next number
		lowerBound := average - stdDev
		upperBound := average + stdDev

		fmt.Printf("Next number range: %.0f %.0f\n", lowerBound, upperBound)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading standard input: %v", err)
	}
}
