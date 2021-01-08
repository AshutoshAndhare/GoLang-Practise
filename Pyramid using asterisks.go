package main

import "fmt"

func NumPyramid() {
	fmt.Println("Enter the height of the pyramid:")
	var height int
	fmt.Scan(&height)
	num := 1
	for i := 0 ; i < height ; i++ {
		
		for j := height + 5 ; j >= i ; j-- {
			fmt.Print("  ")
		} 
		for j:=0 ; j <= i; j++ {
			fmt.Print(num,"   ")
			num++
		}
		num = 1
		fmt.Print("\n\n")
	}
}


func main()  {
	NumPyramid()
}