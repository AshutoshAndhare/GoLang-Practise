package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func indexHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, GO is Neat!") 
}

func newsAggHandler (w http.ResponseWriter, r *http.Request) {            //this is the handle function for the /agg/ page we added
	p := NewsAggPage{Title: "Awesome News Aggregator", News: "some news"} //here we just added some info in the struct 
																		  //which we created to hold info to display on our page
	t, _  := template.ParseFiles("basictemplating.html")				  //this is the line that is using the file 
																		  //"basictemplating.html" as a template on our page
	fmt.Println(t.Execute(w, p))	//in the above line we import the template file and save it in the var t
									//in this line we use that var t to execute the template on our page using the 
									//ResponseWriter w and the struct NewsAggPage using var p
									//And, we put that t.Execute command in Println so that if we make error in the 
									//template file, it will specify the error in our terminal.
}

type NewsAggPage struct { //And this is our struct for trial to store info to display on the /agg/ page
	Title string
	News string
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg", newsAggHandler) //here we are adding a page with add address for trial about templating
	http.ListenAndServe (":8000", nil)
}