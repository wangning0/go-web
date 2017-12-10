package main

import (
	"html/template"
	"log"
	"fmt"
	"os"
	"encoding/json"
	"net/http"
	"strings"
	"errors"
	"data"
)

type Configuration struct {
	Address string
	ReadTimeout int64
	WriteTimeout int64 
	Static string
}

var config Configuration
var logger *log.Logger

// bootstrap log
func p(a ...interface{}) {
	fmt.Println(a)
}

func init() {
	loadConfig()
	file, err := os.OpenFile("chitchat.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("读取日志文件失败", err)
	}
	logger = log.New(file, "INFO", log.Ldate | log.Ltime | log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("读取配置文件失败", err)
		return
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("读取配置文件失败", err)
		return
	}
	file.Close();
}

func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = data.Session{
			Uuid: cookie.Value,
		}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("session 无效")
		}
	}
	return
}

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func info(args ...interface{}) {
	logger.SetPrefix("[INFO] ")
	logger.Println(args)
}

func danger(args ...interface{}) {
	logger.SetPrefix("[ERROR] ")
	logger.Println(args)
}

func warning(args ...interface{}) {
	logger.SetPrefix("[WARNING] ")
	logger.Println(args)
}

func version() string {
	return "0.1"
}