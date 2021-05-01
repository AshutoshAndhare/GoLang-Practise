package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml" 
	"html/template"
)

type SitemapIndex struct { 
	Locations []string `xml:"sitemap>loc"`  
}

type News struct {
	Title []string `xml:"url>news>title"`        
	Keywords []string `xml:"url>news>keywords"`  
	Locations []string `xml:"url>loc"`			 
}

type NewsMap struct { 
	Keyword string    
	Location string   
}

type NewsAggPage struct { 
	Title string
	News map[string]NewsMap
}

func newsAggHandler (w http.ResponseWriter, r *http.Request) {     
	var s SitemapIndex
	var n News
	newsMap := make(map[string]NewsMap) 

	resp, _ := http.Get("https://www.thehindu.com/sitemap/googlenews.xml") 
	bytes, _ := ioutil.ReadAll(resp.Body)  
	resp.Body.Close() 
	xml.Unmarshal(bytes, &s) 
	
	for _, Location := range s.Locations { 
		resp, _ := http.Get(Location)          
		bytes, _ := ioutil.ReadAll(resp.Body)   
		xml.Unmarshal(bytes, &n)           
		resp.Body.Close()
		for idx, _ := range n.Keywords {  
			newsMap[n.Title[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}  
		}
	}

	p := NewsAggPage{Title: "Awesome News Aggregator", News: newsMap} 
	t, _  := template.ParseFiles("NewsAggTemplate.html")				  
	fmt.Println(t.Execute(w, p))	
}

func indexHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, GO is Neat!") 
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg", newsAggHandler) 
	http.ListenAndServe (":8000", nil)
}