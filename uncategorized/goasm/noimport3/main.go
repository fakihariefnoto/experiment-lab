package main

import "fmt"

func Akar(x float64) float64
func bulatBawah(x float64) float64
func bulatAtas(x float64) float64

func percobaan(akar, bulat float64) {
	fmt.Println("Akar dari bilangan 16 adalah ... ", Akar(akar))
	fmt.Println("Pembulatan ke bawah dari 15.56 adalah .. ", bulatBawah(bulat))
	fmt.Println("Pembulatan ke atas dari 15.56 adalah .. ", bulatAtas(bulat))
}

func printPercobaan(akar, bulat float64)

func main() {
	printPercobaan(16, 15.56)
}
