package main

import (
	"fmt"

	"github.com/atselvan/dictionaryapi"
)

func main() {
	fmt.Printf("Starting\n")
	words := []string{"apple", "banana", "cherry", "date", "elderberry", "skjsdf"}
	client := dictionaryapi.NewClient()

	// Lookup the definition for "code"
	for i, word := range words {
		_, err := client.Word.Get(word)
		if err == nil {
			fmt.Printf("%d: %s\n", i, word)
		} else {
			fmt.Printf("%d: %s(%s)\n", i, word, err.Error)
			// log.Fatal(err)
		}
	}
	fmt.Printf("End\n")

}
