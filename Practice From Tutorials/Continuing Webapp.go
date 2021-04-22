package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml" 
)

type SitemapIndex struct { //Here as per the earlier programs of this app, we had 3 structs just for the sitemap index
						   //and the locations tag and converting the locations to string and all.
						   //here we are using just one struct to do all those three.
	Locations []string `xml:"sitemap>loc"`  //here, the Locations element of the struct has a slice that is of the string
											//type which encodes for the loc tag inside the sitemap tag
}

type News struct {
	Title []string `xml:"url>news>title"`        //here we create a news struct so that we can save all the info from
	Keywords []string `xml:"url>news>keywords"`  //the sitemaps inside the sitemaps index i.e. the titles and the keywords
	Locations []string `xml:"url>loc"`			 //and the locations and all
}

type NewsMap struct { // now in the main() we created a map of the news and related info, here we created a struct
	Keyword string    //to save multiple values for a single key in that map
	Location string   //hence, here we have the keywords and the locations and the title would be the key itself
}

func main() {
	var s SitemapIndex
	var n News
	newsMap := make(map[string]NewsMap) //here we created the map for various info of news

	resp, _ := http.Get("https://www.thehindu.com/sitemap/googlenews.xml")  //this is the sitemap index main link
	bytes, _ := ioutil.ReadAll(resp.Body)  
	resp.Body.Close() 
	xml.Unmarshal(bytes, &s) 
	
	for _, Location := range s.Locations { //here we just iterate through the various sitemap links from th sitemap index above
		resp, _ := http.Get(Location)          
		bytes, _ := ioutil.ReadAll(resp.Body)   
		xml.Unmarshal(bytes, &n)           //and store all that info in the News struct for further use
		resp.Body.Close()
		for idx, _ := range n.Keywords {  //here we need the index of the various iterations of the News struct
			newsMap[n.Title[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}  //to use that index to save the specific
																				//values in that index numbers of the map
		}
	}

	for idx, data := range newsMap {  //here we can see the map we created where idx is the key of the map
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
}