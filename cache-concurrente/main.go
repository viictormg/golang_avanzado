package main

import (
	"fmt"
	"sync"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
	lock  sync.Mutex
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
	m.lock.Lock()
	results, exist := m.cache[key]
	m.lock.Unlock()
	if !exist {
		m.lock.Lock()
		results.value, results.err = m.f(key)
		m.cache[key] = results
		m.lock.Unlock()
	}
	return results.value, results.err
}

func GetFibonnaci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {
	cache := NewCache(GetFibonnaci)
	fibo := []int{30, 20, 35, 22}
	var wg sync.WaitGroup
	for _, n := range fibo {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				fmt.Println("error", err)
			}
			fmt.Printf("%d, %v, %d \n", index, time.Since(start), value)
		}(n)
	}
}
