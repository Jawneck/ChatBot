
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
		fmt.Fprintf(w, "Hello, what is your name?")
		return 

	}
	if matched, _:= regexp.MatchString(`(?i).*\bmy name is\b.*`,userinput);matched{
	re := regexp.MustCompile("my name is (.*)") 
	match := re.FindStringSubmatch(userinput)
	topic := match[1]
	fmt.Fprintf(w,"Hello %s", topic)
	return
	}

	strings := []string{
		"I'm not sure what you are talking about.",
		"Could you elaborate on that",
		"I think you are missing the point",
		}
		response := strings[rand.Intn(len(strings))]
		fmt.Fprintf(w, response)
}

func main(){
	rand.Seed(time.Now().UnixNano())		
	
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8080", nil)
}
