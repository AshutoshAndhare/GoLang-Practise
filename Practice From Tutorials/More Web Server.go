package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) { //here we're just looking on how to add html tags in the page
	fmt.Fprintf(w,`<h1>Hey there.</h1> 
<p>Go is fast </p>
<p>...and simple</p>
You can %s add %s.`, "even", "<strong>variables</strong>")
}//and kinda style it a little bit for the inspect source option on the browser

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}