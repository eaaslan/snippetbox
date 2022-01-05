package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tekrar/pkg/forms"
	"tekrar/pkg/models"
)

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{

		Form: forms.New(nil),
	})
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}


	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}



	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}
app.session.Put(r,"flash","Snippet successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}


func (app *application)showSnippet(writer http.ResponseWriter, request *http.Request) {

	id,err :=strconv.Atoi(request.URL.Query().Get(":id"))

	if id<0 || err!=nil {
		app.notFound(writer)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(writer)
		} else {
			app.serverError(writer, err)
		}
		return
	}


app.render(writer,request,"show.page.tmpl",&templateData{
	Snippet:s,
},
	)}

func (app *application)home(writer http.ResponseWriter, request *http.Request) {
	//if request.URL.Path != "/"{
	//
	//	app.notFound(writer)
	//	return
	//}

	q,err := app.snippets.Latest()

	if err!=nil{
		log.Fatalln(err)
	}

app.render(writer,request,"home.page.tmpl",&templateData{Snippets: q})
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {


	app.render(w,r,"signup.page.tmpl",&templateData{Form: forms.New(nil)})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err!=nil{
		app.clientError(w,http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)

	form.Required("name","email","password")
	form.MaxLength("name",255)
	form.MaxLength("email",255)
	form.MatchesPattern("email",forms.EmailRX)

	if !form.Valid(){
		app.render(w,r,"signup.page.tmpl",&templateData{Form: form})
	}

	err = app.users.Insert(form.Get("name"),form.Get("email"),form.Get("password"))

	if err!=nil{
		if errors.Is(err,models.ErrDuplicateEmail){
			form.Errors.Add("email","Address is already in use")
			app.render(w,r,"signup.page.tmpl",&templateData{Form: form})
		}else {
			app.serverError(w,err)
		}
		return
	}
	app.session.Put(r,"flash","Your signup was successful. Please log in.")
	http.Redirect(w,r,"/user/login",http.StatusSeeOther)

	fmt.Fprintln(w, "Create a new user...")
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user login form...")
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login the user...")
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}

