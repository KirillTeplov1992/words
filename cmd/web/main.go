package main

import (
	"database/sql"
	"net/http"
	"fmt"
	"html/template"
	"words/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	topics *mysql.TopicModel
	words *mysql.WordModel
	templateCache map[string]*template.Template
}

func main() {
	dsn := "web:3758@/words?parseTime=true"

	db, err := openDB(dsn)
	if err != nil{
		panic(err)
	}

	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil{
		panic(err)
	}

	app := &application{
		topics: &mysql.TopicModel{DB: db},
		words : &mysql.WordModel{DB: db},
		templateCache : templateCache,	
	}

	fmt.Println("Запуск веб-сервера на http://127.0.0.1:5000")
	err = http.ListenAndServe(":5000", app.routes())
	panic(err) 
}

func openDB(dsn string) (*sql.DB, error){
	db, err:= sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
        return nil, err
	}

	return db, nil
}