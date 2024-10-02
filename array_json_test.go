package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArrayEncode(t *testing.T) {
	customer := Customer{
		FirstName: "Brian",
		LastName:  "Anashari",
		Hobbies:   []string{"Gaming", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Brian","LastName":"Anashari","Age":20,"Student":true,"Hobbies":["Gaming","Coding"]}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Brian",
		LastName:  "Anashari",
		Hobbies:   []string{"Gaming", "Coding"},
		Addresses: []Address{
			{
				Street:     "Jalan Sini",
				Country:    "Indonesia",
				PostalCode: "12345",
			},
			{
				Street:     "Jalan Sini",
				Country:    "Indonesia",
				PostalCode: "67890",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Brian","LastName":"Anashari","Age":0,"Student":false,"Hobbies":["Gaming","Coding"],"Addresses":[{"Street":"Jalan Sini","Country":"Indonesia","PostalCode":"12345"},{"Street":"Jalan Sini","Country":"Indonesia","PostalCode":"67890"}]}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Sini","Country":"Indonesia","PostalCode":"12345"},{"Street":"Jalan Sini","Country":"Indonesia","PostalCode":"67890"}]`

	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}
