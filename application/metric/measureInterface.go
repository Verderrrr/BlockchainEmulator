package metric

type Measure interface {
	OutputRecord() ([]float64, float64)
	OutputMetricName() string
}
