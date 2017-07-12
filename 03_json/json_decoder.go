package main

import (
	"strings"
	"io"
	"fmt"
	"encoding/json"
	"log"
)

// Examples from the standard json package
func main() {
	const jsonStream = `
  		{"Name": "Ed", "Text": "Knock knock."}
  		{"Name": "Sam", "Text": "Who's there?"}
  		{"Name": "Ed", "Text": "Go fmt."}
  		{"Name": "Sam", "Text": "Go fmt who?"}
  		{"Name": "Ed", "Text": "Go fmt yourself!"}
  	`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
	// Output:
	// Ed: Knock knock.
	// Sam: Who's there?
	// Ed: Go fmt.
	// Sam: Go fmt who?
	// Ed: Go fmt yourself!
}
