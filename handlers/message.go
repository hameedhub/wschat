package handlers

import "net/http"

type Message struct{

}

func (m *Message) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
}