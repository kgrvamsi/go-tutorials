package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Humans struct {
	Persons []Person
}

type Person struct {
	Name    string `json:"name"`
	Number  string `json:"number"`
	Created int64  `json:"created"`
}

func main() {

	fmt.Println(time.Now().Unix())

	datas := new(Humans)

	for i := 0; i < 10; i++ {
		//		datas[i] = Person{Name: "Vamsi", Number: "9704087196", Created: time.Now().Unix()}

		datas.Persons = append(datas.Persons, Person{Name: "Vamsi", Number: "9704087196", Created: time.Now().Unix()})
		//		datas.Name = "Vamsi"
		//		datas.Number = "9704087196"
		//		datas.Created = time.Now().Unix()
		//		fmt.Println(datas)
	}

	t := time.Now()

	fmt.Println(t.Format("January 2	8, 2015"))

	jsonString, err := json.Marshal(datas)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(datas)
	fmt.Println(string(jsonString))
}
