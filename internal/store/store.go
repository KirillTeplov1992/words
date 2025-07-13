package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	config *Config
	db *sql.DB
	topicRepository *TopicRepository
	wordRepository *WordRepository

}

func New(config *Config) *Store{
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error{
	db, err:= sql.Open("mysql", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
        return err
	}

	s.db = db

	return nil
}

func (s *Store) Close(){
	s.db.Close()
}

func (s *Store) Topic() *TopicRepository{
	if s.topicRepository != nil{
		return s.topicRepository
	}

	s.topicRepository = &TopicRepository{
		store: s,
	}

	return s.topicRepository
}

func (s *Store) Word() *WordRepository{
	if s.wordRepository != nil{
		return s.wordRepository
	}

	s.wordRepository = &WordRepository{
		store: s,
	}

	return s.wordRepository
}