package handlers

import (
	"net/http"
	"io/ioutil"
)

type Chat struct{

}


func (c *Chat) ServeHTTP(w http.ResponseWriter, r *http.Request){
	ioutil.ReadAll(r.Body)

}