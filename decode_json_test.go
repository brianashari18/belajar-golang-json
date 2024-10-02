package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Brian","LastName":"Anashari","Age":20,"Student":true}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Student)
}
