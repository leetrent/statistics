package statistics

import "math"

func CalcMean(data []float64) float64 {
	var sum float64
	for _, val := range(data) {
		sum += val
	}
	return sum / float64(len(data))
}

func CalcVariance(data []float64) float64 {

	return CalcVarianceUsingMean(CalcMean(data), data)
}

func CalcStandardDeviation( data[] float64) float64 {
	return math.Sqrt(CalcVariance(data))
}

func CalcVarianceUsingMean(mean float64, data []float64) float64 {
	var sqDevFromMean []float64 // squared deviation from mean

	for _, val := range(data) {
		valMinusMean := val - mean
		sqDevFromMean = append(sqDevFromMean, valMinusMean * valMinusMean)
	}

	variance := CalcMean(sqDevFromMean)
	return variance
}

func CalcStandardDeviationUsingVariance(variance float64) float64 {
	return math.Sqrt(variance)
}


