package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

type Result struct {
	XMLName xml.Name `xml:"persons"`
	Persons []Person `xml:"person"`
}
type Person struct {
	Name      string   `xml:"name,attr"`
	Age       int      `xml:"age,attr"`
	Career    string   `xml:"career"`
	Interests []string `xml:"interests>interest"`
}

func main() {
	content, err := ioutil.ReadFile("s.xml")
	if err != nil {
		log.Fatal(err)
	}
	var result Result
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
	log.Println(result.Persons[0].Name)
}
