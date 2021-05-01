package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml" 
	"html/template"
	"sync"
)

var wg sync.WaitGroup

func newsRoutine(c chan News, Location string) { //3.) this takes tha channel name and location as string 
		defer wg.Done() //6.) at last tells the waitgroup that the routine is done
		var n News //4.) creates the var n of News type here instead of the for loop in NewsAggHandle()
		resp, _ := http.Get(Location)          
		bytes, _ := ioutil.ReadAll(resp.Body)   
		xml.Unmarshal(bytes, &n)           
		resp.Body.Close()	
		c <- n //5.) stores value in n in the channel c 
		
}

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

	newsMap := make(map[string]NewsMap) 

	resp, _ := http.Get("https://www.thehindu.com/sitemap/googlenews.xml") 
	bytes, _ := ioutil.ReadAll(resp.Body)  
	resp.Body.Close() 
	xml.Unmarshal(bytes, &s) 
	
	queue := make(chan News, 30) //1.) here we create the channel

	for _, Location := range s.Locations { 
		wg.Add(1)
		go newsRoutine(queue, Location) //2.) start the goroutine with parameters of the channel and Location
	}
	wg.Wait() //7.) waitgroup waits till all the routines are finished and put into channels
	close(queue) //8.) after the routines are complete the channels are closed

	for elem := range queue { //9.) here we iterate over the channels using range since each channels has clusters of n
		for idx, _ := range elem.Keywords {   //10.) here we iterate over all the clusters of n in every channel
		newsMap[elem.Title[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}  //11.) which contain info for the newsMap map
		}
	}

	p := NewsAggPage{Title: "Hindustan Times News Aggregator", News: newsMap} 
	t, _  := template.ParseFiles("NewsAggTemplate.html")				  
	fmt.Println(t.Execute(w, p))	
}

func indexHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<p>Whoa, GO is Neat!</p>
	<p><a href="http://localhost:8000/agg">Click here for News Aggregator Table</a></p>`) 
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg", newsAggHandler) 
	http.ListenAndServe (":8000", nil)
}