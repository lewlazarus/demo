package process

import (
	"fmt"
	"math"
	"sync"

	"demo/data"
	"demo/msg"
	"demo/process/calc"
)

const (
	workersCount = 3
	batchSize    = 15

	OpSum      msg.OperationType = "Sum"
	OpMultiply msg.OperationType = "Multiply"
	OpStdDev   msg.OperationType = "StandardDeviation"
	OpMean     msg.OperationType = "Mean"
	OpVariance msg.OperationType = "Variance"
)

type processFn = func(values []float64) float64
type tempResFn = func(res, batchRes float64) float64

type Processor struct {
	storage    data.StorageInterface
	operations map[msg.OperationType]processFn
}

func NewProcessor(storage data.StorageInterface) *Processor {
	return &Processor{
		storage: storage,
		operations: map[msg.OperationType]processFn{
			OpSum:      opSum,
			OpMultiply: opMultiply,
			OpMean:     opMean,
			OpVariance: opVariance,
			OpStdDev:   opStandardDeviation,
		},
	}
}

func (r *Processor) Process(msg *msg.Request) (float64, float64, error) {
	op, isOk := r.operations[msg.Operation]
	if !isOk {
		return 0, 0, fmt.Errorf("unknown operation '%s'", msg.Operation)
	}

	res := op(msg.Values)

	_ = r.storage.BeginTx()
	prev, _ := r.storage.Get()
	ratio := res / prev

	if res == 0 {
		_ = r.storage.Set(1)
	} else {
		_ = r.storage.Set(res)
	}

	_ = r.storage.EndTx()

	return res, ratio, nil
}

func opSum(values []float64) float64 {
	return process(calc.Sum, sum, 0, values)
}

func opMultiply(values []float64) float64 {
	return process(calc.Multiply, multiply, 1, values)
}

func opMean(values []float64) float64 {
	l := len(values)

	switch l {
	case 0:
		return 0
	case 1:
		return values[0]
	}

	return opSum(values) / float64(l)
}

func opVariance(values []float64) float64 {
	if len(values) <= 1 {
		return 0.0
	}

	mean := opMean(values)
	variance := process(func(v []float64) float64 { return calc.Variance(mean, v) }, sum, 0, values)

	return variance / float64(len(values)-1)
}

func opStandardDeviation(values []float64) float64 {
	return math.Sqrt(opVariance(values))
}

// sum returns result of addition of values a and b
func sum(a, b float64) float64 {
	return a + b
}

// multiply returns result of multiply of values a and b
func multiply(a, b float64) float64 {
	return a * b
}

// process handles the operation calculation process.
// Params:
// processFn - function that actually performs calculation
// tempResFn - function that puts results of processing data batches into a final result
// initVal - initial value for operation
// values - data to process
func process(processFn processFn, tempResFn tempResFn, initVal float64, values []float64) float64 {
	l := len(values)
	if l <= batchSize {
		return processFn(values)
	}

	ch := make(chan [2]int, workersCount)
	chRes := make(chan float64, workersCount)

	wg := sync.WaitGroup{}
	wgRes := sync.WaitGroup{}

	for w := 0; w < workersCount; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				chRes <- processFn(values[v[0]:v[1]])
			}
		}()
	}

	res := initVal
	wgRes.Add(1)
	go func() {
		defer wgRes.Done()
		for batchRes := range chRes {
			res = tempResFn(res, batchRes)
		}
	}()

	posFrom, posTo := 0, 0
	for {
		posTo = posFrom + batchSize
		if posTo > l {
			posTo = l
		}

		ch <- [2]int{posFrom, posTo}

		if posTo >= l {
			break
		}

		posFrom = posTo
	}

	close(ch)
	wg.Wait()

	close(chRes)
	wgRes.Wait()

	return res
}
