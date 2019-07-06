package main

import "fmt"

func Soma(a float64,b float64) float64 {
	return a+b
}
func main(){
	fmt.Printf("%.2f\n", Soma(5,5))
}