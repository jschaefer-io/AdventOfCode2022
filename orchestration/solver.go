package orchestration

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Result struct {
	Error error
	A     string
	B     string
	Time  time.Duration
	TimeA time.Duration
	TimeB time.Duration
}

type Solver interface {
	Setup(data string) error
	A() (string, error)
	B() (string, error)
}

func dispatchDay(solver Solver, path string) Result {
	file, err := os.Open(path)
	if err != nil {
		return Result{
			Error: fmt.Errorf("cannot dispatch from file: %w", err),
		}
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return Result{
			Error: fmt.Errorf("cannot read file content: %w", err),
		}
	}
	dataStr := string(data)
	err = solver.Setup(dataStr)
	if err != nil {
		return Result{
			Error: fmt.Errorf("error on Setup: %w", err),
		}
	}

	startTimeA := time.Now()
	a, err := solver.A()
	timeA := time.Since(startTimeA)
	if err != nil {
		return Result{
			Error: fmt.Errorf("error on A: %w", err),
		}
	}
	startTimeB := time.Now()
	b, err := solver.B()
	timeB := time.Since(startTimeB)
	if err != nil {
		return Result{
			Error: fmt.Errorf("error on B: %w", err),
		}
	}
	return Result{
		A:     a,
		B:     b,
		Time:  time.Since(startTimeA),
		TimeA: timeA,
		TimeB: timeB,
	}
}

func Dispatch(solvers []Solver, path string, from int, to int) (map[int]Result, error) {
	if from-1 < 0 || to-1 >= len(solvers) {
		return nil, fmt.Errorf("dispatch out of range")
	}
	result := make(map[int]Result)
	for i := from; i <= to; i++ {
		fileName := fmt.Sprintf("%s/%02d.txt", path, i)
		result[i] = dispatchDay(solvers[i-1], fileName)
	}
	return result, nil
}
