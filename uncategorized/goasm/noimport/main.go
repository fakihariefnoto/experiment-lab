package main

import "fmt"

func Akar(x float64) float64
func bulatBawah(x float64) float64
func bulatAtas(x float64) float64

func main() {
	fmt.Println("Akar dari bilangan 16 adalah ... ", Akar(16))
	fmt.Println("Pembulatan ke bawah dari 15.56 adalah .. ", bulatBawah(15.56))
	fmt.Println("Pembulatan ke atas dari 15.56 adalah .. ", bulatAtas(15.56))
}
