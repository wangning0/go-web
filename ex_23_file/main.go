package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("winger", 0777)
	os.MkdirAll("winger/1/2/3", 0777)
	err := os.Remove("winger")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("winger")

	userFile := "winger.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.Write([]byte("Just a test\r\n"))
		fout.WriteString("Just a test\r\n")
	}

	fl, error := os.Open(userFile)		
	if error != nil {
		fmt.Println(userFile, error)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}

	e := os.Remove(userFile)

	if e != nil {
		fmt.Println(e)
	}
}