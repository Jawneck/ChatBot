
package main

import (
	"fmt"
	"net/http"
	"regexp"
	"math/rand"
	"time"
)

func userinputhandler(w http.ResponseWriter, r *http.Request){

	userinput :=  r.URL.Query().Get("value")	
	if matched, _:= regexp.MatchString(`(?i).*\bhello\b.*`, userinput);matched{
		fmt.Fprintf(w, "Hello, how are you today?")
		return 
	}

	strings := []string{
		"I'm not sure what you are talking about.",
		"Could you elaborate on that",
		"I think you are missing the point",
		}
		rand.Seed(time.Now().UnixNano())		
		response := strings[rand.Intn(len(strings))]
		fmt.Fprintf(w, response)
}

func main(){
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8080", nil)
}
