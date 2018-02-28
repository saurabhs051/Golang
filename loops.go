package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Loc_array []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

//convertingg derived SitemapIndex to string

func (l Location) String() string { //convert each Location to string	//this fn is called automatically
	return fmt.Sprintf(l.Loc)
}

func main() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// i = 5
	// for { //infinite loop
	// 	fmt.Println("Do Stuff")
	// 	i++
	// 	if i >= 10 {
	// 		break
	// 	}
	// }
	////////////////////////////////////////////
	fmt.Println(" List of sites : ")
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") //Accessing nd getting sitemap
	bytes, _ := ioutil.ReadAll(resp.Body)                                        //Reading the received sitemap as bytes
	var s SitemapIndex
	xml.Unmarshal(bytes, &s) //unmarshalling the received sitemap i.e. getting rid of tags and
	//converting it to string //all Locations stored in array called Loc_array of SitemapIndex variable

	for k, Location := range s.Loc_array { //'range' returns index as well as value at that index
		fmt.Printf("\n%d  %s", k, Location)
	}

}
