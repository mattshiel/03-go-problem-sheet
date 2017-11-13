package main

import (
  "math/rand"
  "time"
  "fmt"
)

// Array of string responses
var responses = []string {
  "I’m not sure what you’re trying to say. Could you explain it to me?",
  "How does that make you feel?",
  "Why do you say that?",
 }

//
func ElizaResponse(input string) string{
  // Adapted info about Intn from https://stackoverflow.com/questions/33994677/pick-a-random-value-from-a-go-slice
  return responses[rand.Intn(len(responses))]
}

func main() {
    // Seed with current time for true random number generation
    rand.Seed(time.Now().UTC().UnixNano())

    // Print input and responses to the console
    fmt.Println("People say I look like both my mother and father.")
    fmt.Println(ElizaResponse("People say I look like both my mother and father."))

    fmt.Println("\nFather was a teacher.")
    fmt.Println(ElizaResponse("Father was a teacher."))

    fmt.Println("\nI was my father’s favourite.")
    fmt.Println(ElizaResponse("I was my father’s favourite."))

    fmt.Println("\nI'm looking forward to the weekend.")
    fmt.Println(ElizaResponse("I'm looking forward to the weekend."))

    fmt.Println("\nMy grandfather was French!")
    fmt.Println(ElizaResponse("My grandfather was French!"))
}
