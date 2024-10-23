package main

import "fmt"

func main() {
	var bilanganBulat, n int
	fmt.Scan(&bilanganBulat, &n)
    hasil := 1

	for i := 1; i <= n; i++ {
		hasil *=  bilanganBulat
	}
	fmt.Print(hasil)
}