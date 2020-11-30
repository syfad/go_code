package main

import "fmt"

type Payer interface {
	pay()
	// duck()
}


type AliPay struct{

}

func (a AliPay) pay() {
	fmt.Println("支付宝pay")
}

type WeixinPay struct{

}


func (w WeixinPay) pay() {
	fmt.Println("微信pay")
}

type YinlianPay struct{

}

func (y YinlianPay) pay() {
	fmt.Println("银联pay")
}

func pay_api(obj Payer)  {
	obj.pay()
}

func main() {

	a:=AliPay{}
    // a.pay()
	//
	w:=WeixinPay{}
    // w.pay()
	//
	y:=YinlianPay{}
	//y.pay()

	pay_api(w)
	pay_api(a)
	pay_api(y)

}
