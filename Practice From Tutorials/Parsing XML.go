package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml" 
)

type Location struct { 
	Loc string `xml:"loc"` //here we create a struct that takes whatever lies between the <loc> and puts that as Loc: whatever
}

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"` //here we create a slice of the Location struct created above and put that as 
										//Locations element of the struct SitemapIndex
}

func (l Location) String() string{ //Here we just created a method for the Location struct just so that whatever we 
									//obtain from the xml file can be changed into a string, else it remains as an obj.
	return fmt.Sprintf(l.Loc) //fmt.Sprintf just turns anything into a string. 
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/arcio/sitemap/master/index/") //Read about this in Accessing the Internet.go
	bytes, _ := ioutil.ReadAll(resp.Body) //this as well 
	resp.Body.Close() //this as well

	var s SitemapIndex //now we create a var s of type SitemapIndex to be used in the main()
	xml.Unmarshal(bytes, &s) //here we take whatever there is in the var bytes and unmarshal it to the address of s
							 //i.e. unmarshal in the SitemapIndex struct.
	fmt.Println(s.Locations) //here we get to see the slice we just created.
}