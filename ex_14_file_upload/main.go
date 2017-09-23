/*
	要使表单能够上传文件，首先第一步就是要添加form的enctype属性，有下面三种情况
	application/x-www-form-urlencoded 表示在发送前编码所有字符(默认)
	multipart/form-data   不对字符编码，在使用包含文件上传控件的表单时，必须使用该值
	text/plain  空格转换为“+”加号，但不对特殊字符编码
*/

package main

import (
	"net/http"
	"time"
	"strconv"
	"fmt"
	"crypto/md5"
	"io"
	"html/template"
	"os"
)

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	if r.Method == "GET" {
		currentTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currentTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		tmpl, _ := template.ParseFiles("form.gtpl")
		tmpl.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8000", nil)
}