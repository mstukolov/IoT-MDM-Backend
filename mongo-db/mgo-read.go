package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

type Model_R struct {
	DataType string
	DataName string
	Attributes []byte
}

type Data struct {
	DataType string `json:"baseType"`
	Atrributes []string `json:"attributes"`
}

type AttrParams struct {
	DisplayName string `json:"displayName"`
	DataType string `json:"dataType"`
}

func main(){
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("mdm").C("model")

	result := Model_R{}
	err = c.Find(bson.M{"datatype": "attrType"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("dataType:", result.DataType)
	fmt.Println("dataName:", result.DataName)
	fmt.Println("attributes:", result.Attributes)


	var t AttrParams
	err2 := json.Unmarshal(result.Attributes, &t)
	if err != nil {
		panic(err2)
	}
	fmt.Println("attributes:", t)

}
