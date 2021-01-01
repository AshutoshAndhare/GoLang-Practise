package main

import "fmt"

func main()  {
	
	nums := []int{2,3,4,6}//here we define and initialize a slice to be used in the range. We could also define an array and rest would have been same
	sum := 0
	for _, num := range nums{//here we want the values of the slice and not the index numbers so we leave a blank '_' before num while declaring
							 // and then we give the range values one by one to var num which goes on adding those to the var sum
		sum += num
	}
	fmt.Println("sum:", sum)//hence at the end we get their total obviously

	for i, num := range nums {// now not always we need the values, sometimes we also need the index numbers of the values,
		if num == 3 {//hence here var i defines the index number of the range, and we tell the if loop to print the index number of
			fmt.Println("index:", i)//only that element of range which has num value of 3
		}
	}

	kvs := map[string]string{"apple":"a", "banana":"b"}
	for k, v := range kvs {//maps also work the same, its just they come in key:value pairs
		fmt.Printf("%s -> %s\n", k,v)//here %s stands for string place holder.
	}

	for k := range kvs {//maps can also show the keys only this way
		fmt.Println("key:", k)
	}

	for i, c := range "go" {//range also shows the unicode code points, where the first value is the starting byte index of the rune
		fmt.Println(i, c)//and second value shows the rune(read below the code for detail) itself
	}
}
/*It is a superset of ASCII and contains all the characters present in the world's writing system including accents 
and other diacritical marks, control codes like tab and carriage return, and assigns each one a standard number called 
a Unicode code point, or in Go language, a rune. The rune type is an alias of int32.*/