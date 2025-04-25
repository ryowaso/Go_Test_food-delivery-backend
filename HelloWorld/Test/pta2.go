package main

import "fmt"

type Account struct {
	name    string
	balance float64
}

func (account *Account) Deposit(amount float64) {
	defer func() {
		fmt.Printf("存款：%s + %.2f\n", account.name, amount)
	}()
	account.balance += amount
}

func (account *Account) Withdraw(amount float64) {
	defer func() {
		if account.balance >= amount {
			fmt.Printf("取款：%s - %.2f\n", account.name, amount)
		}
	}()

	if account.balance < amount {
		fmt.Println("余额不足，无法取款。")
	}
	account.balance -= amount
}

func (account *Account) PrintBalance() {
	fmt.Printf("账户余额：%.2f", account.balance)
}

func main() {
	account := Account{"Alice", 0.0}
	var a, w float64
	fmt.Scanln(&a)
	fmt.Scanln(&w)
	account.Deposit(a)
	account.Withdraw(w)
	account.PrintBalance()

}
