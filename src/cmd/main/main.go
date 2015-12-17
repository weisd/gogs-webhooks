package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Webhook(w http.ResponseWriter, r *http.Request) {
	ret := "ok"
	defer fmt.Fprintf(w, ret)
	err := r.ParseForm()
	if err != nil {
		ret = err.Error()
		return
	}

	fmt.Println(r.Form)

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ret = err.Error()
		return
	}

	fmt.Println(string(data))

	// r.Form.Get(key)

	return
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/hooks", Webhook)
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	log.Fatal(http.ListenAndServe(":8088", mux))
}
