package main

import (
	"fmt"
	"net/http"
)

func (app *application) render (w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		fmt.Println("Шаблон %s не существует!", name)
		return
	}

	err := ts.Execute(w, td)
	if err != nil {
			panic(err)
	}
}