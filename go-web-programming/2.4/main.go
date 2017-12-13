package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	"io"
)

type Post struct {
	Id int `json:"id"`
	Author Author `json:"author"`
	Content string `json:"content"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data: ", err)
		return
	}

	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)

	fmt.Println("==============================================================")


	jsonFile, err = os.Open("post.json")
	if err != nil {
		fmt.Println("error opening json file:",  err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error opening json file:",  err)
			return
		}
		fmt.Println(post)
	}
}