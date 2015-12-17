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
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, ret)
	}()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ret = err.Error()
		return
	}

	var v GogsHookRequest

	err = json.Unmarshal(data, &v)
	if err != nil {
		ret = err.Error()
		return
	}

	fmt.Println(v)

	return
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi-ddd")
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/hooks", Webhook)
	mux.HandleFunc("/", Hello)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
