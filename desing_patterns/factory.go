package main

import "fmt"

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

func (c *Computer) getStock() int {
	return c.stock
}

type Laptop struct {
	Computer
}

func NewLaptop() IProduct {
	return &Laptop{
		Computer{
			name:  "Laptop",
			stock: 25,
		},
	}
}

type Desktop struct {
	Computer
}

func NewDesktop() IProduct {
	return &Desktop{
		Computer{
			name:  "Desktop computer",
			stock: 50,
		},
	}
}

func GetComputerFactory(ComputerType string) (IProduct, error) {
	if ComputerType == "Laptop" {
		return NewLaptop(), nil
	}
	if ComputerType == "Desktop computer" {
		return NewDesktop(), nil
	}
	return nil, fmt.Errorf("invalid type computer")
}

func PrintNameAndStock(p IProduct) {
	fmt.Printf("Product name : %s with stock : %d", p.getName(), p.getStock())
}

func main2() {
	laptop, _ := GetComputerFactory("Laptop")

	PrintNameAndStock(laptop)
}
