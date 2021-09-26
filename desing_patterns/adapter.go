package main

import "fmt"

type Payment interface {
	Pay()
}

type cashPayment struct {
}
type BankPayment struct {
}

func (cashPayment) Pay() {
	fmt.Println("payment using cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

func (b BankPayment) Pay(bankAcount int) {
	fmt.Printf("Pay using bankAcount %d \n", bankAcount)
}

type BanckPaymentAdapter struct {
	BankPayment *BankPayment
	BanckAcount int
}

func (bpa *BanckPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.BanckAcount)
}

func main4() {
	cash := &cashPayment{}
	ProcessPayment(cash)
	bpa := &BanckPaymentAdapter{
		BanckAcount: 5,
		BankPayment: &BankPayment{},
	}

	ProcessPayment(bpa)
}
