// Reading XML sitemap from another site // kinda accessing API
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") //_ is for error variable, as in python
	//Get func will return a response or an error, after trying to read from that URL
	byte, _ := ioutil.ReadAll(response.Body) //response initially in byte form
	str := string(byte)                      //converting to string for our sake
	fmt.Println(str)
	response.Body.Close()

}
