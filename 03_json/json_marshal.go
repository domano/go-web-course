package main

import (
	"fmt"
	"encoding/json"
)

// Examples from the standard json package
func main() {
	var jsonBlob = []byte(`[
  		{"Name": "Platypus", "Order": "Monotremata"},
  		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
  	]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
	// Output:
	// [{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
}
