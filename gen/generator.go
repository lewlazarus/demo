// Helper package for code generation purpose
package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

// genFun a generator function type
type genFun = func(f *os.File, batches []int) error

// main performs code generation.
// CLI arguments: [methodName] [batchSizes...]
// Example: go run generator.go Sum 25 5

func main() {
	var fileName string
	var fn genFun

	args := os.Args[1:]
	op := args[0]

	switch strings.ToLower(op) {
	case "sum":
		fileName = "sum"
		fn = SumGen
	case "multiply":
		fileName = "multiply"
		fn = MultiplyGen
	case "mean":
		fileName = "mean"
		fn = MeanGen
	case "variance":
		fileName = "variance"
		fn = VarianceGen
	case "stddev":
		fileName = "std_dev"
		fn = StdDevGen

	}

	path := "./generated"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	f, err := os.Create(fmt.Sprintf("%s/%s.go", path, fileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	batchMap := make(map[int]bool, 10)
	for _, v := range args[1:] {
		if val, err := strconv.Atoi(v); err != nil {
			panic(err)
		} else {
			if val <= 0 {
				continue
			}

			batchMap[val] = true
		}
	}

	batchMap[1] = true
	batches := make([]int, 0, len(batchMap))
	for val := range batchMap {
		batches = append(batches, val)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(batches)))

	if err := fn(f, batches); err != nil {
		panic(err)
	}

	fmt.Printf("%s generated successfully\n", op)
}

// SumGen generator of Sum method\operation
func SumGen(f *os.File, batches []int) error {

	var temp = template.Must(template.New("").Parse(`// Generated code

package generated

func Sum(values []float64) float64 {
	var res float64 = 0
	{{.Text}}
	return res
}
`))

	s := `var c int

	l := len(values)
	i := 0
`

	for _, v := range batches {
		opsList := make([]string, 0, v)
		for j := 0; j < v; j++ {
			if j == 0 {
				opsList = append(opsList, "values[i]")
			} else {
				opsList = append(opsList, fmt.Sprintf("values[i + %d]", j))
			}

		}

		s += fmt.Sprintf(
			`
	c = (l - i) / %d
	if c > 0 {
		for j := 0; j < c; j++ {
			res += %s
			i += %d
		}
	}
`,
			v, strings.Join(opsList, " + "), v)

	}

	err := temp.Execute(f, struct {
		Text string
	}{
		Text: s,
	})
	if err != nil {
		return err
	}

	return nil
}

// MultiplyGen generator of Multiply method\operation
func MultiplyGen(f *os.File, batches []int) error {

	var temp = template.Must(template.New("").Parse(`// Generated code

package generated

func Multiply(values []float64) float64 {
	var res float64 = 1
	{{.Text}}
	return res
}
`))

	s := `var c int

	l := len(values)
	i := 0
`
	for _, v := range batches {
		opsList := make([]string, 0, v)
		for j := 0; j < v; j++ {
			if j == 0 {
				opsList = append(opsList, "values[i]")
			} else {
				opsList = append(opsList, fmt.Sprintf("values[i + %d]", j))
			}

		}

		s += fmt.Sprintf(
			`
	c = (l - i) / %d
	if c > 0 {
		for j := 0; j < c; j++ {
			res *= %s
			i += %d
		}
	}
`,
			v, strings.Join(opsList, " * "), v)

	}

	err := temp.Execute(f, struct {
		Text string
	}{
		Text: s,
	})
	if err != nil {
		return err
	}

	return nil
}

// MeanGen generator of Mean method\operation
func MeanGen(f *os.File, _ []int) error {
	var temp = template.Must(template.New("").Parse(`// Generated code

package generated

func Mean(values []float64) float64 {
	l := len(values)

	switch l {
	case 0: 
		return 0
	case 1:
		return values[0]
	}

	return Sum(values) / float64(l)
}
`))

	err := temp.Execute(f, struct{}{})
	if err != nil {
		return err
	}

	return nil
}

// VarianceGen generator of Variance method\operation
func VarianceGen(f *os.File, batches []int) error {

	var temp = template.Must(template.New("").Parse(`// Generated code

package generated

import "math"

func Variance(values []float64) float64 {
	l := len(values) 
	if l <= 1 {
		return 0.0
	}

	var res float64 = 0
	{{.Text}}
	return res / float64(l) 
}
`))

	s := `var c int

	m := Mean(values)
	i := 0
`

	for _, v := range batches {
		opsList := make([]string, 0, v)
		for j := 0; j < v; j++ {
			if j == 0 {
				opsList = append(opsList, "math.Pow(values[i]-m, 2)")
			} else {
				opsList = append(opsList, fmt.Sprintf("math.Pow(values[i + %d]-m, 2)", j))
			}

		}

		s += fmt.Sprintf(
			`
	c = (l - i) / %d
	if c > 0 {
		for j := 0; j < c; j++ {
			res += %s
			i += %d
		}
	}
`,
			v, strings.Join(opsList, " + "), v)

	}

	err := temp.Execute(f, struct {
		Text string
	}{
		Text: s,
	})
	if err != nil {
		return err
	}

	return nil
}

// StdDevGen generator of StandardDeviation method\operation
func StdDevGen(f *os.File, _ []int) error {
	var temp = template.Must(template.New("").Parse(`// Generated code

package generated

import "math"

func StandardDeviation(values []float64) float64 {
	return math.Sqrt(Variance(values))
}
`))

	err := temp.Execute(f, struct{}{})
	if err != nil {
		return err
	}

	return nil
}
