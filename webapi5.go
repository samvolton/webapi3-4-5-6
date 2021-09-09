package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title           string
	Author          string
	Header          string
	PageDescription string
	Content         string
	URI             string
}

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	//String birleştirme işlemi
	var builder bytes.Buffer
	builder.WriteString("Something of something;\n")
	builder.WriteString("31pages\n")
	builder.WriteString("ISSS: 69-31-Adult\n")
	builder.WriteString("Price: 0\n")
	builder.WriteString("Size: 31 x 31\n")
	builder.WriteString("2.Baski\n")

	uri := "www.wologame.com/313131/"

	page := Page{
		Title:           "Adult: qwdqwfqwfqwf",
		Author:          "XXX YYY",
		Header:          "WEGERHERHRWGHERHREWH",
		PageDescription: "sdbewefgewgw",
		Content:         builder.String(),
		URI:             "http://" + uri}

	t, _ := template.ParseFiles("page1.html")
	t.Execute(w, page)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
