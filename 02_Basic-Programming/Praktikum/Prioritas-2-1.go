package main

import "fmt"

func main(){
	//Segitiga Asterik
	var num int
	fmt.Printf("Masukkan Panjang Segitiga: ")
	fmt.Scan(&num)
	for i:=1; i<=num;i++{
		for ii:=1; ii<=num-i;ii++{
			fmt.Printf(" ")
		}
		for j:=1; j<=i;j++{
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}

	
}