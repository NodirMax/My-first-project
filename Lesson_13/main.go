package main

import (
	"errors"
	"fmt"
)

type User interface {
	Balance_cheking() (float64, error)
	Paying(float64) (float64, error)
	Transfer(QR, float64) (float64, float64, error)
}

type QR struct {
	id       uint
	UserName string
	Balance  float64
}

func (u QR) Balance_cheking() (float64, error) {
	if u.Balance < 0.1 {
		err := errors.New("Service is not available")
		return 0, err
	} else {
		u.Balance -= 0.1
		return u.Balance, nil
	}
}

func (u QR) Paying(product float64) (float64, error) {
	if u.Balance < product {
		err := errors.New("You haven't enough money!")
		return u.Balance, err
	} else {
		u.Balance -= product
		return u.Balance, nil
	}
}

func (u QR) Transfer(v QR, count float64) (float64, float64, error) {
	if count > v.Balance {
		err := errors.New("You haven't enough money")
		return u.Balance, v.Balance, err
	} else {
		u.Balance += count
		v.Balance -= count
		return u.Balance, v.Balance, nil
	}
}

func Rezult(u User) {
	balance, err := u.Balance_cheking()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Your balance now: ", balance)
	}
}

func Pay(u User) {
	var banana float64 = 5.02 // price of product
	balance, err := u.Paying(banana)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your balance now: ", balance)
	}
}

func TNF(u User) {
	Andry := QR{id: 2, UserName: "Andry", Balance: 1200}
	value := 400.00
	userBalance, _, _ := u.Transfer(Andry, value)
	fmt.Print("Transfer completed successfully")
	fmt.Println("Your balance now:", userBalance)

}
func main() {
	Jimmy := QR{id: 1, UserName: "Jimmy", Balance: 900}
	Rezult(Jimmy)
	Pay(Jimmy)
	TNF(Jimmy)
}
