package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	Foo string `json:"foo"`
}
func main(){
	text := []byte(`{"foo":"bar"}`)
	var t T
	err := json.Unmarshal(text, &t)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}
