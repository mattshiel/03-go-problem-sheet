package main

import (
	"fmt"
	"math/rand"
	"regexp" // Import Golang's built-in support for regukar expressions
	"strings"
	"time"
)

// Array of string responses
var responses = []string{
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
}

func applyPronouns(input string) string {
	// Fields splits the string input around one or more consecutive white space characters
	sentence := strings.Fields(input)

	// Map of string to string key-value pairs, to substitute pronouns
	// List of substitutions adapted from https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/
	pronouns := map[string]string{
		"i":      "you",
		"was":    "were",
		"i'd":    "you would",
		"i've":   "you have",
		"i'll":   "you will",
		"my":     "your",
		"are":    "am",
		"you've": "I have",
		"you'll": "I will",
		"your":   "my",
		"yours":  "mine",
		"you":    "I",
		"me":     "you",
	}

	// Loop through slice checking at each index if the word is equal to a key in the map
	// If the word is equal to the map value, change the word to equal the value
	// Found how to iterate over a map from https://gobyexample.com/range
	for index, word := range sentence {
		if value, ok := pronouns[strings.ToLower(word)]; ok {
				sentence[index] = value;
		}
	}

	return strings.Join(sentence, " ")
}

func ElizaResponse(input string) string {
	// regex expression to search case-insensitively for the string "father"
	// Learned about regex from https://github.com/StefanSchroeder/Golang-Regex-Tutorial/blob/master/01-chapter1.markdown
	match := regexp.MustCompile(`(?i)\bfather\b`)

	// If the match is false then the input is checked to see if it begins with "I am"
	if match.MatchString(input) {
		return "Why don’t you tell me more about your father?"
	} else {
		// Use the "|" operator to match groups so as to define seperate submatches
		// Escape full stops and question marks
		match := regexp.MustCompile(`(?i)^(I am|I'm|im) ([^\.\?!]*)`)

		if match.MatchString(input) {
			// FindStringSubmatch returns a slice of strings
			// Position 2 will find the subexpressions as defined by the submatch description in the regex
			captured := match.FindStringSubmatch(input)[2]
			captured = applyPronouns(captured)
			return fmt.Sprintf("How do you know you are %s?", captured)
		}
	}
	// Otherwise we return a random response
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

	fmt.Println("\nI am happy.")
	fmt.Println(ElizaResponse("I am happy."))

	fmt.Println("\nI am not happy with your responses.")
	fmt.Println(ElizaResponse("I am not happy with your responses."))

	fmt.Println("\nI am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))

	fmt.Println("\nI am supposed to just take what you’re saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you’re saying at face value?"))
}
