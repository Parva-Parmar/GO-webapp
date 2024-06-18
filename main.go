package main

import(
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Hello World")
}

func main() {
	http.HandleFunc("/",helloWorldPage)
	http.ListenAndServe("",nil)
}