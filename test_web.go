// package main

// import (
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", index)
// 	log.Printf("start ... \n")
// 	http.ListenAndServe(":8080", nil)
// }

// func index(w http.ResponseWriter, r *http.Request) {
// 	content, _ := ioutil.ReadFile("./index.html")
// 	w.Write(content)
// }
