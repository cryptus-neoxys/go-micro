package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func main() {
    hello := func (w http.ResponseWriter, r *http.Request)  {
        log.Println("Hello World")
    }
    bye := func (w http.ResponseWriter, r *http.Request) {
        log.Println("Bbye World")
    }

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        log.Println("Hello, World")

        d, err := ioutil.ReadAll(r.Body)
        if(err != nil) {
            http.Error(w, "Ooopsie", http.StatusBadRequest)
            return
            // <- Same as above -> 👆🏻
            // w.WriteHeader(http.StatusBadRequest)
            // w.Write([]byte ("Ooopsie"))
            // return
        }
        fmt.Fprintf(w, "Data: %s", d)
    })
        
    http.HandleFunc("/hello", hello)
        
    http.HandleFunc("/bye", bye)
    
    http.ListenAndServe(":42069", nil)
}