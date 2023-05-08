package msg

type OperationType string

// Request represents request data structure
type Request struct {
	Operation OperationType `json:"operation"`
	Values    []float64     `json:"values"`
}
