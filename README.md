# ChatBot #
# Author: Danielis Joniskis, G00333859 #
I am a third year student in Galway-Mayo Institute of Technology. I am studying Software Development.

This is the project assigned to us for Data Representation and Querying. The goal was to create a chatbot web application in Go. The user can access the chatbot via their browser. They can input text and will recieve a reply based on their choice of words. Example: "Hello my name is Danielis", the chatbot will reply to this with "Hello Danielis!". Initially I had the input box at the top of the page but chose to move it to the footer as this would allow me to scroll the page and make the whole messaging visual more realistic. I ran into problems with appending in the output box at first, instead I appended the input and output to a list which was much easier to manipulate. I ran into some issues with the footer not being sticky but got it fixed with a bit of trial and error. Overall I am quite happy with the outcome.

1. Clone the repository
2. cd into the folder
3. Run the program "go run eliza.go"
4. View in "http://localhost:8080/"