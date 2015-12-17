package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Webhook(w http.ResponseWriter, r *http.Request) {
	ret := "ok"
	var err error
	defer func(w http.ResponseWriter, ret string, err error) {
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, ret)
	}(w, ret, err)

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ret = err.Error()
		return
	}

	var v GogsHookRequest

	err = json.Unmarshal(data, v)
	if err != nil {
		ret = err.Error()
		return
	}

	fmt.Println(GogsHookRequest)

	return
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi-d")
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

	log.Fatal(http.ListenAndServe(":8080", mux))
}
