package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct {
	fName   string
	sName   string
	age     int
	thought []byte
}

type Abbas Person

func (a Abbas) GetFirstName() string {
	return a.fName
}

func (a Abbas) GetSecondName() string {
	return a.sName
}

func (a Abbas) GetAge() int {
	return a.age
}

func (a Abbas) GetThought() []byte {
	return a.thought
}

func (a Abbas) String() string {
	return fmt.Sprintf("Abbas's thought is: %s", string(a.thought))
}

//This writer function, write something to Abbas's thought
func (a *Abbas) Write(p []byte) (n int, err error) {
	a.thought = make([]byte, len(p))
	for i, v := range p {
		a.thought[i] = v
	}

	return len(a.thought), nil
}

func main() {
	//opening a file
	f, err := os.Open("thought.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//initializing a pointer to Abbas type
	p := &Abbas{} //same as new(Abbas)

	//copy from io.Reader to io.Writer
	//in our case, io.Reader: file & io.Writer: Abbas(p)
	//and go compiler, automatically finds Write function
	b, err := io.Copy(p, f)
	if err != nil {
		log.Fatal((err))
	} else if b == 0 {
		log.Fatal("nothing was written")
	}

	fmt.Println(p)
}
