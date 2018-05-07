package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	DataType string `json:"dataType"`
	DisplayName string `json:"displayName"`
}
type Model struct {
	Data Data `json:"model"`
}

func main(){

	text := []byte(`{"model":{"dataType":"storeType","displayName":"Типы магазинов"}}`)

	var t Model
	err := json.Unmarshal(text, &t)

	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}
