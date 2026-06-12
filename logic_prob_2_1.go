package main

import "fmt"

func isPrima(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func logic_prob_2_1() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	if isPrima(n) {
		fmt.Printf("%d adalah bilangan prima\n", n)
	} else {
		fmt.Printf("%d bukan bilangan prima\n", n)
	}
}
