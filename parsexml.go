package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"` //for  <sitemap> tag
}

type Location struct {
	Loc string `xml:"loc"` //for <loc> tag
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc) //Sprintf converts it to a specified format and returns string
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	var s SitemapIndex
	xml.Unmarshal(bytes, &s) //unmarshals the sitemapindex
	fmt.Println(s.Locations)

}
