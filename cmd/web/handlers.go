package main

import (
	"net/http"
	"fmt"
	"strconv"
)

func (app *application) home (w http.ResponseWriter, r *http.Request) {
	topics := app.topics.GetTopics()
	
	app.render(w, r, "home.page.tmpl", &templateData{
		Topics : topics,
	})
}

func (app *application) create (w http.ResponseWriter, r *http.Request){
	app.render(w, r, "create.page.tmpl", nil)
}

func (app *application) addTopic (w http.ResponseWriter, r *http.Request){
	name := r.FormValue("name")

	app.topics.Insert(name)

	http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
}

func (app *application) openTopic (w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	topic, err := app.topics.GetContentOfTopicById(id)
	if err != nil {
		panic(err)
	}

	app.render(w, r, "open_topic.page.tmpl", &templateData{
		Topic : topic,
	})	
}

func (app *application) openFormAddWord(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	topic, err := app.topics.GetContentOfTopicById(id)

	app.render(w, r, "add_word.page.tmpl", &templateData{
		Topic : topic,
	})
	
}

func (app *application) addWord (w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	word := r.FormValue("word")
	translation := r.FormValue("translation")

	app.topics.AddWordToBase(id, word, translation)

	http.Redirect(w, r, fmt.Sprintf("/topic?id=%d",id), http.StatusSeeOther)
}

func (app *application) editTopic (w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	topic, err := app.topics.GetTopic(id)
	if err != nil {
		panic(err)
	}

	app.render(w, r, "edit_topic.page.tmpl", &templateData{
		TopicName : topic,
	})	
}

func (app *application) updateTopic (w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	app.topics.UpdateTopic(id, name)

	http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
}

func (app *application) openEditWordPage (w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	word, err := app.words.GetWord(id)
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	app.render(w, r, "edit_word.page.tmpl", &templateData{
		Word : word,
	})
}

func (app *application) editWord (w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	word := r.FormValue("word")
	translation := r.FormValue("translation")

	app.words.UpdateWord(id, word, translation)

	http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
}