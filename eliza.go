
package main

import (
	"fmt"
	"net/http"
	"regexp"
	"math/rand"
	"time"
)

//A function which handles the users input(string) and manipulates it to give a response.
func userinputhandler(w http.ResponseWriter, r *http.Request){

	//The user input is gathered from the user input field in the form residing in index.html
	userinput :=  r.URL.Query().Get("value")
	//Using regexp to match the string hello so that the chat bot can give a response.
	//Here we match the string case-insensitively.
	if matched, _:= regexp.MatchString(`(?i).*\bHello\b|\bHi\b|\bHey\b.*`, userinput);matched{
		fmt.Fprintf(w, "Hello, what is your name?")
		return 
	}
	//Here we are taking in every string after the string 'my name is'.
	//This allows us to then give a response using the name the user has provided. 
	if matched, _:= regexp.MatchString(`(?i).*\bMy name is\b.*`,userinput);matched{
		re := regexp.MustCompile("(?i)My name is (.*)") 
		match := re.FindStringSubmatch(userinput)
		topic := match[1]
		fmt.Fprintf(w,"Hello %s!", topic)
		return
	}
	//This match allows the user to ask who/what the chat bot is.
	if matched, _:= regexp.MatchString(`(?i).*\bWho are you|What are you|Whats your name\b.*`,userinput);matched{
		fmt.Fprintf(w, "My name is Eliza.\n")
		fmt.Fprintf(w, "I am a chat bot.\n")
		fmt.Fprintf(w, "What are you?")		
		return 
	}
	//Here the user can respond to the question by stating I am a/I'm a ____.
	//Everything after 'a' is taken in as user input and manipulated back to the user. 
	if matched, _:= regexp.MatchString(`(?i).*\bI am a|I'm a\b.*`,userinput);matched{
		re := regexp.MustCompile("[I am|I'm] a (.*)") 
		match := re.FindStringSubmatch(userinput)
		topic := match[1]
		fmt.Fprintf(w, "Oh, do you like being a %s?", topic)		
		return 
	}
	//In this we match two different string that appear anywhere in the input line.
	//Here the user can respond to the question "tell me why like/love", by using the word because
	if matched, _:= regexp.MatchString(`(?i).*?\bI like\b.*?\bbecause\b.*?|I love\b.*?\bbecause\b.*?`,userinput);matched{
		fmt.Fprintf(w, "Oh thats good.\n")
		fmt.Fprintf(w, "Have you got anything else to tell me?")		
		return 
	}
	//In this we match two different string that appear anywhere in the input line.
	//Here the user can respond to the question "tell me why you don't like/hate", by using the word because
	if matched, _:= regexp.MatchString(`(?i).*?\bI don't like\b.*?\bbecause\b.*?|I hate\b.*?\bbecause\b.*?`,userinput);matched{
		fmt.Fprintf(w, "Thats a shame.\n")
		fmt.Fprintf(w, "Have you got anything else to tell me?")		
		return 
	}
	//The user can state that they like/love something.
	if matched, _:= regexp.MatchString(`(?i).*\bI like|I love\b.*`,userinput);matched{
		re := regexp.MustCompile("[I like|I love] (.*)") 
		match := re.FindStringSubmatch(userinput)
		topic := match[1]
		fmt.Fprintf(w, "Tell me why you %s?", topic)		
		return 
	}
	//The user can state that they don't like/hate something.
	if matched, _:= regexp.MatchString(`(?i).*\bI don't like|I hate\b.*`,userinput);matched{
		re := regexp.MustCompile("[I don't like|I hate] (.*)") 
		match := re.FindStringSubmatch(userinput)
		topic := match[1]
		fmt.Fprintf(w, "Tell me why you %s?", topic)		
		return 
	}
	//The chat bot giving a response to the user saying no/nope
	if matched, _:= regexp.MatchString(`(?i).*No|Nope.*`, userinput);matched{
		fmt.Fprintf(w, "Why so negative?")
		return 
	}
	//Here the chat bot will respond with the user saying goodbye.
	if matched, _:= regexp.MatchString(`(?i).*\bGoodbye|Bye\b.*`,userinput);matched{
		fmt.Fprintf(w, "Farewell friend, until next time.\n")	
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
