package main

func login(w http.ResponseWriter, req *http.Request, ctx *Context) (err error) {
  //grab the username and password from the form
  username, password := req.FormValue("username"), req.FormValue("password")

  //log in the user
  user, err := Login(ctx, username, password)

  //what to do now? if there was an error we want to present the form again
  //with some error message.

  //where do we store the user if the login was valid?

  //answer: sessions!
  _ = user
  return
}

var login = parseTemplate(
  "templates/_base.html",
  "templates/login.html",
)

func loginForm(w http.ResponseWriter, req *http.Request, ctx *Context) (err error) {
  err = login.Execute(w, nil)
  return
}

func loginForm(w http.ResponseWriter, req *http.Request, ctx *Context) (err error) {
  return T("login.html").Execute(w, nil)
}
