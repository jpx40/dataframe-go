// Copyright 2018-20 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package hw

// See: http://www.itl.nist.gov/div898/handbook/pmc/section4/pmc435.htm
func initialTrend(y []float64, period int) float64 {

	var sum float64
	sum = 0.0

	for i := 0; i < period; i++ {
		sum += (y[period+i] - y[i]) / float64(period)
	}

	return sum / float64(period)
}

// See: http://www.itl.nist.gov/div898/handbook/pmc/section4/pmc435.htm
func initialSeasonalComponents(y []float64, period int, tsType Method) []float64 {

	nSeasons := len(y) / period

	seasonalAverage := make([]float64, nSeasons)
	seasonalIndices := make([]float64, period)

	// computing seasonal averages
	for i := 0; i < nSeasons; i++ {
		for j := 0; j < period; j++ {
			seasonalAverage[i] += y[(i*period)+j]
		}
		seasonalAverage[i] /= float64(period)
	}

	// Calculating initial Seasonal component values

	for i := 0; i < period; i++ {
		for j := 0; j < nSeasons; j++ {
			if tsType == Multiplicative {
				// Multiplcative seasonal component
				seasonalIndices[i] += y[(j*period)+i] / seasonalAverage[j]
			} else {
				// Additive seasonal component
				seasonalIndices[i] += y[(j*period)+i] - seasonalAverage[j]
			}

		}
		seasonalIndices[i] /= float64(nSeasons)
	}

	return seasonalIndices
}
