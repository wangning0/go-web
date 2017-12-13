package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	post := Post {
		Id :1,
		Content: "hello world",
		Author: Author {
			Id: 2,
			Name: "winger",
		},
		Comments: []Comment{
			Comment {
				Id: 3,
				Content: "wang",
				Author: "ning",
			},
			Comment {
				Id: 4,
				Content: "ning",
				Author: "nick",
			},
		},
	}

	output, err := json.MarshalIndent(post, "", "\t")
	if err != nil {
		fmt.Println("error marshal error: ", err)
		return
	}
	err = ioutil.WriteFile("post.json", output, 0644)
	if err != nil {
		fmt.Println("error writing json :", err)
		return
	}
	

	jsonFile, err := os.Create("post1.json")
	if err != nil {
		fmt.Println(err, "error")
		return
	}
	defer jsonFile.Close()
	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println(err, "error ")
		return
	}
}