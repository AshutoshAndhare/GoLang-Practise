package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/arcio/sitemap/master/index/") 
	//This above one just says the url that we want the info you have stored on the specific page. Hence, we http.Get()
	//and the Store that in the var resp.
	bytes, _ := ioutil.ReadAll(resp.Body)
	//in the above one, we tell the var resp that we need your Body and then tell tell ioutil to just read all that there is
	//in resp.Body and store that in the var bytes.
	stringBody := string(bytes)
	//in the var bytes, we got all the info in the form of bytes. Hence, we convert that info to string 
	//and store that string info in var stringBody
	fmt.Println(stringBody)
	//here in the above one, we just print all that stringBody on the terminal
	resp.Body.Close()
	//here above, we just close the Get() request we started in the first line of code.
}