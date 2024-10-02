package golangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("customer_output.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Brian",
		LastName:  "Anashari",
	}

	encoder.Encode(customer)
}
