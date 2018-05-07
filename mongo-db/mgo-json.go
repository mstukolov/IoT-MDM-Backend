package main

import (
	"log"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type Model struct {
	DataType string
	DataName string
	//Attributes []byte
	Attributes string
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

	//attr := []byte(`{"model":{"dataType":"storeType","displayName":"Типы магазинов"}}`)

	/*j_chain := []byte(`{
      "baseType": "chain",
      "attributes" : [
        {"guid": {"displayName": "Koд записи", "type": "int"}},
        {"name": {"displayName": "Название", "type": "string"}}
      ]
    }`)*/

	//j_attr := []byte(`{"displayName": "Koд записи", "type": "int"}`)
	j_attr := `{"displayName": "Koд записи", "type": "int"}`

	model := Model{"attrType", "chain", j_attr}

	err = c.Insert(&model)
	if err != nil {
		log.Fatal(err)
	}


}
