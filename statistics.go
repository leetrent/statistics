package statistics

import (
	"errors"
	"math"
)

func CalcMean(data []float64) float64 {
	var sum float64
	for _, val := range data {
		sum += val
	}
	return sum / float64(len(data))
}

func CalcVariance(data []float64) float64 {

	return CalcVarianceUsingMean(CalcMean(data), data)
}

func CalcStandardDeviation(data []float64) float64 {
	return math.Sqrt(CalcVariance(data))
}

func CalcVarianceUsingMean(mean float64, data []float64) float64 {
	var sqDevFromMean []float64 // squared deviation from mean

	for _, val := range data {
		valMinusMean := val - mean
		sqDevFromMean = append(sqDevFromMean, valMinusMean*valMinusMean)
	}

	variance := CalcMean(sqDevFromMean)
	return variance
}

func CalcStandardDeviationUsingVariance(variance float64) float64 {
	return math.Sqrt(variance)
}

func CalcStandardError(popStdDev float64, sampleSize int64) (float64, error) {
	if sampleSize == 0 {
		return 0.0, errors.New("sampleSize cannot by zero")
	}
	return popStdDev / math.Sqrt(float64(sampleSize)), nil
}
