package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/",healthz)

	err := http.ListenAndServe(":80",nil)
	if err != nil {
		fmt.Println(err)
	}
}
func healthz(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	fmt.Fprintln(w,"request header参数:")
	for key,values := range header {
		for _,value := range values {
			fmt.Fprintln(w,key,"=",value)
		}
	}

	version := os.Getenv("VERSION")
	fmt.Fprintln(w,"\n环境变量VERSION is :",version)

	addr := strings.Split(r.RemoteAddr,":")
	fmt.Println("remote ip:",addr[0])

	var statusCode int = 200
	w.WriteHeader(statusCode)
	fmt.Println("statusCode : ",statusCode)
}
