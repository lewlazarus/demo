package msg

type OperationType string

const (
	OpSum      OperationType = "Sum"
	OpMultiply OperationType = "Multiply"
	OpStdDev   OperationType = "StandardDeviation"
	OpMean     OperationType = "Mean"
	OpVariance OperationType = "Variance"
)

// Request represents request data structure
type Request struct {
	Operation OperationType `json:"operation"`
	Values    []float64     `json:"values"`
}
