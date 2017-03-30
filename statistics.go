package statistics

import (
	"errors"
	"fmt"
	"math"
	"sort"
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
		return math.NaN(), errors.New("sampleSize cannot by zero")
	}
	return popStdDev / math.Sqrt(float64(sampleSize)), nil
}

func CalcMedian(data []float64) (float64, error) {

	var median float64

	// make sure that slice is not empty
	count := len(data)
	if count == 0 {
		return math.NaN(), errors.New("Length of data cannot by zero")
	}

	// make the sure the slice is sorted in ascending order
	if sort.Float64sAreSorted(data) == false {
		sort.Float64s(data)
	}

	// get the midpoint index
	mid := float64(count) / float64(2)

	// is the sample size count even or odd?
	if count%2 == 0 {
		lbIdx := int(mid) - 1 // lower-bound index
		ubIdx := lbIdx + 1    // upper-bound index
		lbVal := data[lbIdx]  // lower-bound value
		ubVal := data[ubIdx]  // upper-bound value

		median = (lbVal + ubVal) / 2
	} else {
		median = data[int(mid)]
	}

	return median, nil
}

func test1() {
	data := []float64{26.1, 25.6, 25.7, 25.2, 25.0}
	fmt.Println(data)
	med, err := CalcMedian(data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(med)
	}
}

func test2() {
	data := []float64{24.7, 25.0, 25.2, 25.6, 25.7, 26.1}
	fmt.Println(data)
	med, err := CalcMedian(data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(med)
	}
}
func main() {
	//http://www.statcan.gc.ca/edu/power-pouvoir/ch11/median-mediane/5214872-eng.htm
	//https://play.golang.org/p/OWbFm_XBoE
	test1()
	fmt.Println("")
	test2()
}
