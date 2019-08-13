package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running now")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var res string
	fmt.Println("Receiving Request from ", r.RemoteAddr)
	w.WriteHeader(200)
	host, err := os.Hostname()
	if err != nil {
		fmt.Println("Error in getting host name")
	}
	res = fmt.Sprintf("You have hit %s\n", host)
	w.Write([]byte(res))
}
