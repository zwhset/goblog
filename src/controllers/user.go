package controllers

import (
	"fmt"
	"html/template"
	"log"
	"modules"
	"net/http"
)

func Login(res http.ResponseWriter, req *http.Request) {
	result := Result{}
	// Post Action
	if req.Method == "POST" {
		result.IsPost = true

		req.ParseForm() // init parse Form
		u := modules.User{
			Username: req.FormValue("username"),
			Password: req.FormValue("password"),
		}
		fmt.Println(u.Username, u.Password)
		_, err := u.Login()

		if err != nil {
			result.Code = false
			result.Message = err.Error()
		} else {
			// set cookie
			cookie := http.Cookie{
				Name:   "session-id",
				Value:  "session-value",
				Path:   "/",
				MaxAge: 3600,
			}

			http.SetCookie(res, cookie)

			http.Redirect(res, req, "/", 301)
			return
		}

	}

	tpl, err := template.ParseFiles("views/login.html", "views/header.tpl", "views/footer.tpl")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(res, result)
	if err != nil {
		log.Fatal(err)
	}

}
