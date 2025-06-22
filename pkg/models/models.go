package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Topic struct {
	ID int
	Name string
}

type Word struct{
	ID int
	Word string
	Translation string
}

type TopicContent struct{
	TopicID int
	TopicName string
	Words []*Word
}