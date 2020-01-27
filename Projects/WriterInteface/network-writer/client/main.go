package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/*
JSON MESSAGE:

{
	"thought": "I'm busy"
}

*/

var jsonMessage = `{"thought": "I'm busy"}`

//postFromFile sends a post request to a server, taking a json body from a file
func postFromFile() error {
	f, err := os.Open("thought.json")
	if err != nil {
		return errors.New("postFromFile: " + err.Error())
	}

	_, err = http.Post("http://127.0.0.1:5050/thought", "application/json", f)
	if err != nil {
		return errors.New("postFromFile: " + err.Error())
	}

	return nil
}

//postFromBytesBuffer sends a post request to a server, taking a json body from bytes buff qer
func postFromBytesBuffer() error {
	buf := &bytes.Buffer{}

	err := json.NewEncoder(buf).Encode(jsonMessage)
	if err != nil {
		return errors.New("postFromBytesBuffer: " + err.Error())
	}

	_, err = http.Post("http://127.0.0.1:5050/thought", "application/json", buf)
	if err != nil {
		return errors.New("postFromBytesBuffer: " + err.Error())
	}
	
	return nil
}

func getResponse() error {
	resp, err := http.Get("http://127.0.0.1:5050/thought")
	if err != nil {
		return errors.New("getResponse: " + err.Error())
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("getResponse: " + err.Error())
	}

	fmt.Printf("Status is: %s\n", resp.Status)
	fmt.Printf("Response is: %s\n", string(b))

	return nil
}

func main() {
	var err error

	// err = postFromBytesBuffer()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = postFromFile()
	if err != nil {
		log.Fatal(err)
	}

	err = getResponse()
	if err != nil {
		log.Fatal(err)
	}
}
