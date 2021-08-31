package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Post struct {
	Id      int       `json:"id"`
	Content string    `json:"content"`
	Author  Author    `json:"author"`
	Comment []Comment `json:"comments"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening json file:", err)
		return
	}
	defer jsonFile.Close()

	// jsonData, err := ioutil.ReadAll(jsonFile)
	// if err != nil {
	// 	fmt.Println("Error reading json file:", err)
	// 	return
	// }

	// post := Post{}

	// err = json.Unmarshal(jsonData, &post)
	// fmt.Println(post)

	post1 := Post{}
	decoder := json.NewDecoder(jsonFile)
	for {
		err = decoder.Decode(&post1)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoder json file:", err)
			return
		}

	}
	fmt.Println(post1)

	post2 := Post{
		Id:      2,
		Content: "GO",
		Author: Author{
			Id:   3,
			Name: "lisi",
		},
		Comment: []Comment{
			{
				Id:      4,
				Content: "0831",
				Author:  "wangwu",
			},
			{
				Id:      5,
				Content: "0901",
				Author:  "zhaoliu",
			},
		},
	}
	// fmt.Println(post2)
	// post2Data, err := json.MarshalIndent(&post2, "", "\t\t")
	// post2Data, err := json.Marshal(&post2)
	// if err != nil {
	// 	fmt.Println("Error marshal json post2:", err)
	// 	return
	// }
	// err = ioutil.WriteFile("post2.json", post2Data, 0644)
	// if err != nil {
	// 	fmt.Println("Error writing json:", err)
	// 	return
	// }
	post3Data, err := os.Create("post3.json")
	if err != nil {
		fmt.Println("Error creating json:", err)
		return
	}
	encoder := json.NewEncoder(post3Data)
	err = encoder.Encode(&post2)
	if err != nil {
		fmt.Println("Error encoding json:", err)
		return
	}
}
