package main

import "fmt"

func main()  {
	s := make([]string, 3)//declaring a slice using a make statement, just the syntax is changed as against to arrays where the
						  // number of '[]' defines how many dimentions in the slice, and number of elements are declared after the declaration of the data types
	fmt.Println("emp:", s)

	s[0] = "a"//now this part works totally like arrays
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))//and this as well

	s = append(s, "d")//now this is not the one in arrays, here, unlike arrays, we can add more elements then we declared in the make statement by using the append statement.
	s = append(s, "e", "f")//like here we first defined s to be a slice with 3 elements, but then using append we added 3 more.
	fmt.Println("adp:", s)

	c := make([]string, len(s))//now in slice we can also copy an array using the copy statement.
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]//here we use the slice statement that has syntax of [low:high], in this case which sliced the elements from 3 to 5 including 3 and 5.
	fmt.Println("sl1:", l)//which means the input will be [c d e]

	l = s[:5]//here it'll show the first 5 elements i.e. index no. = 0,1,2,3,4.
	fmt.Println("sl2:", l)

	l = s[2:]//here it'll skip the first two elements i.e. it'll show all emements excluding index no. 0,1 
	fmt.Println("sl3:", l)

	t:= []string{"g","h","i"}//this how we declare and initialize the slice in a single line
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)//here we declare a two dimentional slice but the number of elements in the inner slices is not declared 
	for i := 0;i < 3; i++ {//so the number of elements in the inner slices can vary and like here we can declare the number of elements later
		innerLen := i + 1
		twoD[i]=make([]int, innerLen)//we declare the number of elements in inner slices here.
		for j := 0;j < innerLen;j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoD)
}