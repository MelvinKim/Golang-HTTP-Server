package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello world!")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ooops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Hello %s", data)
		// log.Printf("Data %s\n", data)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":9090", nil)
	/*
		constructs an http server and registers a default handler
		if no second parameter is specified , it uses the default server mux?
	*/
}

// :9090 -> binds all requests on all IPs on port 9090

/*
- handleFunc is a convenient method on the http package
- registers a function to a Path, on the default server Mux
- default serve Mux is an http handler
*/

/*
- Reading and writing to requests
- we use http.Request and http.ResponseWriter
- you can use ResponseWrite to :
1. write Headers
2. write Response Body
3. write write Status Codes
- the request contains:
1. the path used
2. the response body that was passed from the client
3. the HTTP methood used eg GET, POST , PUT, PATCH , DELETE
4. the HTTP version
*/
