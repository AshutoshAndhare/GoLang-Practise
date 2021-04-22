package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml" 
)

type Location struct { 
	Loc string `xml:"loc"` 
}

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`  
}

func (l Location) String() string{ 
	return fmt.Sprintf(l.Loc) 
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/arcio/sitemap/master/index/") 
	bytes, _ := ioutil.ReadAll(resp.Body)  
	resp.Body.Close() 
	var s SitemapIndex 
	xml.Unmarshal(bytes, &s) 

	for _, Location := range s.Locations { //so this is the loop where we just iterate through the various Locations 
											//elements of the SitemapIndex struct and save the values in the var Location
		fmt.Printf("\n%s", Location) //Here we print that var Location on  a new line everytime 
	}
}