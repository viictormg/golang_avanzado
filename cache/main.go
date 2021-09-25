package main

import (
	"fmt"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Celdas de memoria
type Memory struct {
	f     Function
	cache map[int]FunctionResult
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	result, exist := m.cache[key]

	if !exist {
		result.value, result.err = m.f(key)
		m.cache[key] = result
	}
	return result.value, result.err
}

func GetFibonnaci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {
	cache := NewCache(GetFibonnaci)
	// fibo := []int{77, 60, 50, 77, 30, 35}
	fibo := []int{50, 7, 4, 50}

	for _, n := range fibo {
		start := time.Now()
		value, err := cache.Get(n)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d, %s, %d\n", n, time.Since(start), value)
	}

}
