package app

import "net/http"

func (app *Application) configureRouter() {
	app.router.HandleFunc("/", app.home)
	app.router.HandleFunc("/create", app.create)
	app.router.HandleFunc("/add_topic", app.addTopic)
	app.router.HandleFunc("/topic", app.openTopic)
	app.router.HandleFunc("/topic/{id}/add_word_form", app.openFormAddWord)
	app.router.HandleFunc("/topic/{id}/add_word", app.addWord)
	app.router.HandleFunc("/edit_topic", app.editTopic)
	app.router.HandleFunc("/update_topic", app.updateTopic)
	app.router.HandleFunc("/word", app.openEditWordPage)
	app.router.HandleFunc("/edit_word", app.editWord)

	//подключаю CSS стили
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    app.router.Handle("/static/", http.StripPrefix("/static", fileServer))
}