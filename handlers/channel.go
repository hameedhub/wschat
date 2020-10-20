package handlers

import (
	"net/http"
	"io/ioutil"
)

type Channel struct{

}


func (c *Channel) ServeHTTP(w http.ResponseWriter, r *http.Request){
	rs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
	w.Write([]byte(rs))
}