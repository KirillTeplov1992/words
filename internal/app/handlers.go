package app

import (
	"fmt"
	"net/http"
	"strconv"
	"words/internal/models"
)

func (app *Application) home (w http.ResponseWriter, r *http.Request) {
	topics := app.store.Topic().GetTopics()
	
	app.render(w, r, "home.page.tmpl", &templateData{
		Topics : topics,
	})

}

func (app *Application) create (w http.ResponseWriter, r *http.Request){
	t := models.Topic{
		Name: "",
	}
	app.render(w, r, "create.page.tmpl", &templateData{
		TopicName: &t,
		Errors: []string{},
	})
}

func (app *Application) addTopic (w http.ResponseWriter, r *http.Request){
	name := r.FormValue("name")

	var errList []string

	if name == ""{
		t := models.Topic{
			Name: name,
		}

		errList = append(errList, "Введите название темы")
		app.render(w, r, "create.page.tmpl", &templateData{
			TopicName: &t,
			Errors: errList,
	})
	} else{
		app.store.Topic().Insert(name)
		http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
	}
}

func (app *Application) openTopic (w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	topic, err := app.store.Topic().GetContentOfTopicById(id)
	if err != nil {
		panic(err)
	}

	app.render(w, r, "open_topic.page.tmpl", &templateData{
		Topic : topic,
	})	
}

func (app *Application) openFormAddWord(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	topic, err := app.store.Topic().GetContentOfTopicById(id)

	wordModel := models.Word{
		Word: "",
		Translation: "",
	}

	app.render(w, r, "add_word.page.tmpl", &templateData{
		Topic : topic,
		Word: &wordModel,
	})
	
}

func (app *Application) addWord (w http.ResponseWriter, r *http.Request){
	var errList []string
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	word := r.FormValue("word")
	translation := r.FormValue("translation")

	wordModel := models.Word{
		Word: word,
		Translation: translation,
	}

	topic, err := app.store.Topic().GetContentOfTopicById(id)
	if err != nil{
		panic(err)
	}

	if word == ""{
		errList = append(errList, "Введите слово")
	}
	if translation == ""{
		errList = append(errList, "Введите перевод")
	}

	if len(errList) > 0{
		app.render(w, r, "add_word.page.tmpl", &templateData{
		Topic : topic,
		Word: &wordModel,
		Errors: errList,
	})
	} else {
		app.store.Topic().AddWordToBase(id, word, translation)
		http.Redirect(w, r, fmt.Sprintf("/topic?id=%d",id), http.StatusSeeOther)
	}	
}

func (app *Application) editTopic (w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	topic, err := app.store.Topic().GetTopic(id)
	if err != nil {
		panic(err)
	}

	app.render(w, r, "edit_topic.page.tmpl", &templateData{
		TopicName : topic,
	})	
}

func (app *Application) updateTopic (w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	app.store.Topic().UpdateTopic(id, name)

	http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
}

func (app *Application) openEditWordPage (w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	word, err := app.store.Word().GetWord(id)
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	app.render(w, r, "edit_word.page.tmpl", &templateData{
		Word : word,
	})
}

func (app *Application) editWord (w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w,r)
		return
	}

	word := r.FormValue("word")
	translation := r.FormValue("translation")

	app.store.Word().UpdateWord(id, word, translation)

	http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
}