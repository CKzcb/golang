// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", index)
// 	log.Printf("start ... \n")
// 	http.ListenAndServe(":8080", nil)
// }

// func index(w http.ResponseWriter, r *http.Request) {
// 	log.Println(r)
// 	fmt.Fprintf(w, "hi~")
// }
