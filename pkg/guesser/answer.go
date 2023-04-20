package guesser

import "math"

// Answer is the result of a guesslang model execution, also providing whether or not the answer is considered reliable.
type Answer struct {
	Predictions []Prediction
	Reliable    bool
}

// Prediction is a single language prediction, including the confidence and the language.
type Prediction struct {
	Confidence float64
	Language   string
}

// IsReliable returns true if the prediction is considered "reliable".
// It is considered reliable if the probability of the predicted language is higher than 2 standard deviations from the mean.
// Original: https://github.com/yoeo/guesslang/blob/42ec63776777e1bdce2d72f51710c6634e36e00c/guesslang/guess.py#L157-L165
func IsReliable(confidences []float64) bool {
	mean := mean(confidences)
	stdev := stdev(confidences, mean)
	threshold := mean + 2*stdev
	predictedLanguageProbability := max(confidences)
	return predictedLanguageProbability > threshold
}

// stdev returns the standard deviation of a slice of float64.
func stdev(values []float64, mean float64) float64 {
	var sum float64
	for _, v := range values {
		sum += math.Pow(v-mean, 2)
	}
	return math.Sqrt(sum / float64(len(values)))
}

// mean returns the mean of a slice of float64.
func mean(values []float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

// max returns the maximum value of a slice of float64.
func max(values []float64) float64 {
	var maxVal = math.Inf(-1)
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
