package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

// go test -v
//to test our program
//execute in test mode

var testOk = `1 
2
3 
4 
5`

var testOkResult = `1 
2
3 
4 
5`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err != nil {
		t.Errorf("test for Ok Failed")
	}
	if out.String() != testOkResult {
		t.Errorf("test for OK failed - result dodnt match")
	}
}
