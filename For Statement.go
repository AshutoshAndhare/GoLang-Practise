package main

import "fmt"

func main() {

    i := 1
    for i <= 3 {//this is the basic for statement where variable is declared out of loop and the loop limit is defined with the loop
        fmt.Println(i)
        i = i + 1//and the increment is shown within the loop
    }

    for j := 7; j <= 9; j++ {//this is classic for loop, where variable is declared in scope and validity and increment is declared with the loop
        fmt.Println(j)
    }

    for {//this loop doesn't need a condition to run and it runs repeatedly till we break the loop with a break statement.
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {//classic for loop but contains continue statement, which skips the loop one step(next interation of the loop) for the statement part and ahead 
        if n%2 == 0 {
            continue//mostly this statement is used with conditionals like if and switch.
        }
        fmt.Println(n)
    }
}