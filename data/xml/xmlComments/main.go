package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Post struct {
	XMLName xml.Name  `xml:"post"`
	Id      string    `xml:"id,attr"`
	Content string    `xml:"content"`
	Author  Author    `xml:"author"`
	Xml     string    `xml:",innerxml"`
	Comment []Comment `xml:"comments>comment"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening Xml file:", err)
		return
	}
	defer xmlFile.Close()

	// xmlData, err := ioutil.ReadAll(xmlFile)
	// if err != nil {
	// 	fmt.Println("Error reading Xml file:", err)
	// 	return
	// }

	// var post Post
	// xml.Unmarshal(xmlData, &post)
	// fmt.Println(post)

	decode := xml.NewDecoder(xmlFile)

	for {
		token, err := decode.Token()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}

		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decode.DecodeElement(&comment, &se)
				fmt.Println(comment)
			}
		}

	}
}
