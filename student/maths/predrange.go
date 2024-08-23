package maths

// Predict the next number range based on the mean and standard deviation of the last N numbers
func PredictRange(numbers []float64) (float64, float64) {
	m := Mean(numbers)
	sd := StandardDeviation(numbers)
	// We can use a prediction range of mean ± 1.5*standard deviation
	return m - 2*sd, m + 2*sd
}