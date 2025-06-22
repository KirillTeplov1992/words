package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/create", app.create)
	mux.HandleFunc("/add_topic", app.addTopic)
	mux.HandleFunc("/topic", app.openTopic)
	mux.HandleFunc("/topic/{id}/add_word_form", app.openFormAddWord)
	mux.HandleFunc("/topic/{id}/add_word", app.addWord)
	mux.HandleFunc("/edit_topic", app.editTopic)
	mux.HandleFunc("/update_topic", app.updateTopic)

	//подключаю CSS стили
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}