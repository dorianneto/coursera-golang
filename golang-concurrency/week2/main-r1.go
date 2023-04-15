package main

import (
	"fmt"
	"time"
)

func main() {
	balance := 0

	/**
	This block just for loop 10 times to add 100 and minus 100 from balance to simulate a bank
	transaction, where you need enough balance to withdraw
	*/
	for i := 0; i < 10; i += 1 {
		go func() {
			deposit := 100
			fmt.Printf("Step 1: Deposit $%v\n", deposit)
			balance += deposit
		}()

		go func() {
			withdraw := 100
			fmt.Printf("Step 2: Withdraw $%v\n", withdraw)
			if balance < withdraw {
				fmt.Printf("not enough fund to withdraw $%v, balance is %v\n", withdraw, balance)
				return
			}
			balance -= withdraw
			fmt.Printf("Withdraw succeed, balance is %v\n", balance)
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("final balance is $%v\n", balance)
	/**
	Race condition occur because the Deposit and Withdraw read and write value of balance concurrently
	without synchronization therefore we could see final balance end up not becoming $0 and so sometimes
	there is not enough balance to withdraw.
	*/
}
