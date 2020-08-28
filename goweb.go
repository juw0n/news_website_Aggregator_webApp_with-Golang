package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat!")
	fmt.Fprintf(w, "<p>Go is fast!</p>")
	fmt.Fprintf(w, "<p>...and simple!</p>")
	fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>variables</strong>")
	fmt.Fprintf(w, `
		<h6>You can even do ...</h6>
		<h5>multiple lines ...</h5>
		<h4>in one %s</h4>`, "formatted print")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expert web design by Juwon")
}

func main() {
	http.HandleFunc("/", indexHandler) //the HandleFunc handles functions that will coorespond to paths which in this case is "/".
	http.HandleFunc("/about/", aboutHandler)
	http.ListenAndServe(":8000", nil) //listen on what port?   ... can serve on tls with ListenAndServeTLS ... need something in server args, we'll put nil for now
}
