package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	
)

func main() {
	res, _ := http.Get("http://www.google.com/")
	
	page, error := ioutil.ReadAll(res.Body)

	res.Body.Close()

	fmt.Printf("%s", page)
	fmt.Printf("%s", error)

}