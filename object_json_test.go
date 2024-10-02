package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street, Country, PostalCode string
}

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Student   bool
	Hobbies   []string
	Addresses []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Brian",
		LastName:  "Anashari",
		Age:       20,
		Student:   true,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
