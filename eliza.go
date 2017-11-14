
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
	fmt.Fprintf(w,"Hello %s!", topic)
	return
	}
	if matched, _:= regexp.MatchString(`(?i).*\bwho are you|what are you\b.*`,userinput);matched{
		fmt.Fprintf(w, "I am a chat bot.\n")
		fmt.Fprintf(w, "What are you?")		
		return 
	}
	if matched, _:= regexp.MatchString(`(?i).*\bI am a|I'm a\b.*`,userinput);matched{
		re := regexp.MustCompile("[I am|I'm] a (.*)") 
		match := re.FindStringSubmatch(userinput)
		topic := match[1]
		fmt.Fprintf(w, "Oh, do you like being a %s?", topic)		
		return 
	}
	if matched, _:= regexp.MatchString(`(?i).*\bI like|I love\b.*`,userinput);matched{
		re := regexp.MustCompile("[I like|I love] (.*)") 
		match := re.FindStringSubmatch(userinput)
		topic := match[1]
		fmt.Fprintf(w, "Tell me why you %s?", topic)		
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
