package main

import (
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	fName   string
	sName   string
	age     int
	Thought []byte `json:"thought"`
}

type Abbas Person

//A instance of Abbas's type
var A = &Abbas{}

func (a Abbas) String() string {
	return fmt.Sprintf("Abbas's thought is: %s", string(a.Thought))
}

//This writer function, write something to Abbas's thought
func (a *Abbas) Write(p []byte) (n int, err error) {
	a.Thought = make([]byte, len(p))
	for i, v := range p {
		a.Thought[i] = v
	}

	return len(a.Thought), nil
}

func thoughtHandler(w http.ResponseWriter, r *http.Request) {
	//check method of request
	switch r.Method {

	case http.MethodGet:
		w.Header().Add("Status", "200")

		//write to response our message
		//check if a type has a string method and then call it
		//and write string method return to w (response)
		fmt.Fprint(w, A)

	case http.MethodPost:
		w.Header().Add("Status", "200")

		//write to our type's field: Thought
		//request's body
		//call Write function
		fmt.Fprint(A, r.Body)
	}
}

func main() {
	//default server
	srv := http.Server{
		Addr:    "127.0.0.1:5050",
		Handler: http.DefaultServeMux,
	}

	//give a route to multiplexer
	http.HandleFunc("/thought", thoughtHandler)

	log.Fatal(srv.ListenAndServe())
}
