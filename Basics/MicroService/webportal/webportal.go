package webportal

import "net/http"

import "fmt"

// To run our Portal that it is addressing on it
func RunWebPortal(link string) error {
	http.HandleFunc("/", roothandler)
	return http.ListenAndServe(link, nil)
}

//In order to do not wtite too many arguments we simply create fuction named roothandler
func roothandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Portal %s", r.RemoteAddr)
}
