package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Define simple responses for specific inputs
func getResponse(input string) string {
    input = strings.ToLower(strings.TrimSpace(input))

    switch input {
    case "hello", "hi":
        return "Hello! How can I help you today?"
    case "how are you?":
        return "I'm just a bunch of code, but I'm doing fine!"
    case "what is your name?":
        return "I'm a simple Go chatbot."
    case "bye", "exit":
        return "Goodbye! Have a great day!"
    default:
        return "Sorry, I don't understand that. Can you try asking something else?"
    }
}

func main() {
    fmt.Println("Welcome to GoBot! Type something to begin (type 'exit' to quit).")

    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("You: ")
        if !scanner.Scan() {
            break
        }
        input := scanner.Text()
        if strings.ToLower(strings.TrimSpace(input)) == "exit" {
            fmt.Println("GoBot: Goodbye!")
            break
        }

        response := getResponse(input)
        fmt.Println("GoBot:", response)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Error reading input:", err)
    }
}

