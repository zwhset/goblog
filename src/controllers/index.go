package controllers

import (
	"html/template"
	"log"
	"modules"
	"net/http"
)

type Person struct {
	UserName string
	Age      int
}

func Index(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("views/index.html", "views/header.tpl")
	if err != nil {
		log.Fatal(err)
	}

	t := modules.Topic{}
	data := t.ViewAll()
	err = tpl.Execute(res, data)
	if err != nil {
		log.Fatal(err)
	}
}
