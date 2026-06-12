package main

import "fmt"

func main() {
	susu := 3.97
	roti := 2.17
	banana := 0.99
	apple := 0.89

	var qtySusu, qtyRoti, qtyBanana, qtyApple int
	fmt.Print("Jumlah susu: ")
	fmt.Scan(&qtySusu)
	fmt.Print("Jumlah roti: ")
	fmt.Scan(&qtyRoti)
	fmt.Print("Jumlah banana: ")
	fmt.Scan(&qtyBanana)
	fmt.Print("Jumlah apple: ")
	fmt.Scan(&qtyApple)

	// Hitung dengan discount
	totalSusu := float64(qtySusu/2)*5.00 + float64(qtySusu%2)*susu
	totalRoti := float64(qtyRoti/3)*6.00 + float64(qtyRoti%3)*roti
	totalBanana := float64(qtyBanana) * banana
	totalApple := float64(qtyApple) * apple

	total := totalSusu + totalRoti + totalBanana + totalApple

	fmt.Printf("\nTotal: $%.2f\n", total)
}
