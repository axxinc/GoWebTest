package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	file, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file::", err)
		return
	}
	defer file.Close()

	xmlData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)

	post1 := Post{
		Id:      "2",
		Content: "Hello World",
		Author: Author{
			Id:   "3",
			Name: "Zhang San",
		},
	}
	output, err := xml.Marshal(&post1)
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}

	err = ioutil.WriteFile("post2.xml", output, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// add \t begins on a new indented line
	output2, err := xml.MarshalIndent(&post1, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}

	err = ioutil.WriteFile("post3.xml", output2, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// add xml.Header
	output3, err := xml.MarshalIndent(&post1, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}

	err = ioutil.WriteFile("post4.xml", []byte(xml.Header+string(output3)), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Encoder
	post5file, err := os.Create("post5.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}
	encoder := xml.NewEncoder(post5file)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post1)
	if err != nil {
		fmt.Println("Error encoding XML to file:", err)
		return
	}
}
