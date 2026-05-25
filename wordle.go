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
	var greenValues []string = []string{"", "", "", "", "T"}
	var yellowValues []string = []string{"", "", "I", "S", ""}
	var grayValues string = "alernoychump"

	fmt.Printf("%q\n", greenValues)
	fmt.Printf("%q\n", yellowValues)
	fmt.Printf("%q\n", grayValues)
	fmt.Printf("End\n")

}
