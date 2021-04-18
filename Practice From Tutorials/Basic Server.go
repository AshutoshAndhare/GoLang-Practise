package main

import (
	"fmt"
	"net/http"
)
//first go to the func main and start reading comments from there, then come to the below function
func indexHandler (w http.ResponseWriter, r *http.Request) { // Here, w and r are var names which can be anything,
															  //while their types i.e. http.ResponseWriter and http.Request
															  //are custom types which we are not supposed to know rn the 
															  //only thing we know about them is that it has something to 
															  //do with structs, i.e. http like the var and ResponseWriter 
															  //being the element of the struct.
	fmt.Fprintf(w, "Whoa, GO is Neat!") //now here, Fprintf will just use our ResponseWriter w to write whatever we want.
}

func main() {
	/*For any web server or web page, there is a need of a handler who, if you type a 
	certain address, decides what to do with that address and what to show the user for entering the perticular address.
	Hence, it HANDLES the situation after the user enters a certain address. The function http.HandleFunc() is a handler
	here.*/
	http.HandleFunc("/", indexHandler) //So, this is a handler function , where the first parameter is the path(address), 
										//And the 2nd parameter is the function, which should start when we enter the path

	http.ListenAndServe (":8000", nil)  //And this is a function which creates a local server on port :8000 and the 
										//second parameter is unknnown to me so the tutorial guy said to enter nil.

}