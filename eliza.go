
package main

import (
	"fmt"
	"net/http"
//	"regexp"
	"math/rand"
	"time"
)

func userinputhandler(w http.ResponseWriter, r *http.Request) {

	strings := []string{
		"I'm not sure what you are talking about.",
		"Could you elaborate on that",
		"I think you are missing the point",
		}
		rand.Seed(time.Now().UnixNano())
		response := strings[rand.Intn(len(strings))]

	userinput :=  r.URL.Query().Get("value")
	if(userinput == "Hello" || "hello"){
	fmt.Fprintf(w, "%s there senor",userinput)
	}else{
		fmt.Fprintf(w, response)
	}

}
func main(){
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8080", nil)
}
