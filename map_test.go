package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id":"P001","name":"Brian Anashari","balance":100000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)
	fmt.Println(result)
}

func TestMapDecode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "Brian Anashari",
		"price": 100000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
