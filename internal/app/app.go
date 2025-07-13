package app

import (
	"html/template"
	"net/http"
	"words/internal/store"

	"github.com/sirupsen/logrus"
)

type Application struct {
	config *Config
	logger *logrus.Logger
	router *http.ServeMux
	store *store.Store
	templateCache map[string]*template.Template
}

func New (config *Config) *Application{
	return &Application{
		config : config,
		logger: logrus.New(),
		router: http.NewServeMux(),
	}
}

func (app *Application) Start() error{
	if err := app.configureLogger(); err != nil{
		return err
	}

	app.configureRouter()

	if err := app.configureStore(); err != nil{
		return err
	}

	if err := app.configureTemplates(); err != nil{
		return err
	}

	app.logger.Info("Запуск веб-сервера на http://127.0.0.1:5000")
	return http.ListenAndServe(app.config.BindAddr, app.router)
}

func (app *Application) configureLogger() error{
	level, err := logrus.ParseLevel(app.config.LogLevel)
	if err != nil{
		return err
	}

	app.logger.SetLevel(level)

	return nil
}

func (app *Application) configureStore() error{
	st := store.New(app.config.Store)
	if err := st.Open(); err != nil{
		return err
	}

	app.store = st

	return nil
}

func (app *Application) configureTemplates() error{
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil{
		return err
	}

	app.templateCache = templateCache

	return nil
}

func (app *Application) render (w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.logger.Printf("Шаблон %s не существует!", name)
		return
	}

	err := ts.Execute(w, td)
	if err != nil {
			panic(err)
	}
}