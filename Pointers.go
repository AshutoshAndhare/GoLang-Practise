package main

import "fmt"

func zeroval(ival int)  {
	ival = 0
}

func zeropts(iptr *int)  {
	*iptr = 0
}

func main()  {
	i := 1
	fmt.Println("intial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeropts(&i)
	fmt.Println("zeropts:", i)

	fmt.Println("pointer:", &i)
	
}