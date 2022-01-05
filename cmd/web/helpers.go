package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

func (app *application)addDefaultData(td *templateData, r *http.Request)*templateData  {

	if td == nil{
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	td.Flash = app.session.PopString(r,"flash")

	return td

}
func (app *application)render(w http.ResponseWriter,r *http.Request,name string,td *templateData)  {
	ts,ok :=app.templateCache[name]
	if !ok{
		log.Fatalln(ok)
	}
	buf := new(bytes.Buffer)
	err :=ts.Execute(w,app.addDefaultData(td,r))

	if err!=nil{
		log.Fatalln(err)
	}
	buf.WriteTo(w)
}

func (app *application)serverError(w http.ResponseWriter,err error) {
	trace := fmt.Sprintf("%s \n %s",err.Error(),debug.Stack())
	http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	app.errorLog.Output(2,trace)
}

func (app *application) clientError(w http.ResponseWriter,status int)  {

	http.Error(w,http.StatusText(status),status)
}
func (app *application)notFound(w http.ResponseWriter)  {
	app.clientError(w,http.StatusNotFound)

}

