package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (pay payment) makepayment(amount float32) {
	// make instance of phone pay first
	// mypayment := phonepay{}
	// mypayment.pay(amount)

	//no need to create instance og google pay we have the interface
	pay.gateway.pay(21)
}

type phonepay struct {
}

func (phone phonepay) pay(amount float32) {
	fmt.Println("making payment using phone pay",amount)
}

type googlepay struct {
}

func (google googlepay) pay(amount float32) {
	fmt.Println("making payment using google pay",amount)
}

func main() {
	// make instance of payments first
	// mypayments := payment{}
	// mypayments.makepayment(230000)

	googly := googlepay{}
	newpayment := payment{
		gateway: googly,
	}
	newpayment.makepayment(100)

}
