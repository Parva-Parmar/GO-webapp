package main

import(
	"fmt"
	"net/http"
)

// func handleFunction(w http.ResponseWriter, r *http.Request){
// 	switch r.URL.Path{
// 	case "/":
// 		fmt.Fprint(w,"Hello World")
// 	case "/ninja":
// 		fmt.Fprint(w,"Wallace")
// 	default:
// 		fmt.Fprint(w,"Big Fat Error!")
// 	}
// 	fmt.Printf("Handling function with %s request\n",r.Method)
// }

func htmlVsPlain(w http.ResponseWriter, r *http.Request){
	fmt.Println("htmlVsPlain")
	fmt.Fprint(w,"<h1>Hello World</h1>")
}

func main() {
	http.HandleFunc("/hp",htmlVsPlain)
	http.ListenAndServe("",nil)
}