package main

import (
	"bufio"
	"fmt"
	"guess/maths"
	"os"
	"strconv"
)

func main() {
	const windowSize = 5 // Number of last numbers to consider for prediction
	var data []float64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		data = append(data, value)

		if len(data) >= windowSize {
			data = data[len(data)-windowSize:]
		}

		if len(data) > 1 {
			lower, upper := maths.PredictRange(data)
			fmt.Printf("%d %d\n", int(lower), int(upper))
		}
	}
}