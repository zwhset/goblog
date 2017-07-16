package controllers

import (
	"html/template"
	"log"
	"modules"
	"net/http"
)

func SetSession(res http.ResponseWriter, name, value, path string, maxAge int, httpOnly bool) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		MaxAge:   maxAge,
		HttpOnly: httpOnly,
	}
	http.SetCookie(res, &cookie)
}

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
		_, err := u.Login()

		if err != nil {
			result.Code = false
			result.Message = err.Error()
		} else {
			// set cookie
			SetSession(res, u.Username, "session-value", "/", 3600, true)
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

func Register(res http.ResponseWriter, req *http.Request) {
	result := Result{}
	// Post Action

	if req.Method == "POST" {
		result.IsPost = true

		req.ParseForm() // init parse Form
		u := modules.User{
			Username:   req.FormValue("username"),
			Password:   req.FormValue("password"),
			RePassword: req.FormValue("repassword"),
		}
		_, err := u.Register()

		if err != nil {
			result.Code = false
			result.Message = err.Error()
		} else { // 注册session
			// set cookie
			SetSession(res, u.Username, "session-value", "/", 3600, true)
			http.Redirect(res, req, "/login", 301)
			return
		}

	}

	tpl, err := template.ParseFiles("views/register.html", "views/header.tpl", "views/footer.tpl")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(res, result)
	if err != nil {
		log.Fatal(err)
	}

}
