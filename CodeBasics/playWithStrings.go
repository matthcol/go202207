package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Play With Strings")
	sentence := "The quick brown fox jumped over the lazy dog."
	fmt.Println(sentence)
	fmt.Println(strings.ToLower(sentence))
	fmt.Println(strings.ToUpper(sentence))
	b := strings.Contains(sentence, "fox")
	fmt.Println("fox is there:", b)
	b = strings.ContainsAny(sentence, "wzx")
	fmt.Println("any letter wzx is there:", b)
	b = strings.ContainsRune(sentence, '福')
	fmt.Println("福 letter is there:", b)
	pos := strings.Index(sentence, "fox")
	fmt.Println("fox is at pos:", pos)
	pos = strings.Index(sentence, "chicken")
	fmt.Println("chicken is at pos:", pos)
	fmt.Println(strings.HasPrefix(sentence, "The"))
	fmt.Println(strings.HasSuffix(sentence, "dog."))
	words := strings.Fields(sentence)
	newSentence := strings.Join(words, ", ")
	fmt.Println(newSentence)
	ok, error := regexp.MatchString("[a-z]+", sentence)
	fmt.Println(ok, error)
	pattern, _ := regexp.Compile("[a-z]+")
	words = pattern.FindAllString(sentence, -1)
	fmt.Println(words)
}
