package mysql

import (
	"database/sql"
	"words/pkg/models"
)

type TopicModel struct{
	DB *sql.DB
}

func (m *TopicModel) GetTopics() ([]*models.Topic) {
	stmt := `SELECT
				id,
				name
			FROM
				topic`
	res, err := m.DB.Query(stmt)
		if err != nil {
			panic(err)
		}

	var topics []*models.Topic

	for res.Next() {
		topic := &models.Topic{}
		err = res.Scan(&topic.ID, &topic.Name)
		if err != nil{
			panic(err)
		}

		topics = append(topics, topic)
	}

	return topics
}

func (m *TopicModel) Insert (name string) (int, error){
	stmt := `INSERT INTO topic (
				name)
			VALUES (
				?)`
	result, err := m.DB.Exec(stmt, name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err !=nil {
		return 0, err
	}

	return int(id), nil
}

func (m *TopicModel) GetContentOfTopicById (id int)(*models.TopicContent, error){
	topic, err := m.GetTopic(id)

	stmt := `SELECT 
				id,
				word,
				translation
			FROM 
				words
			WHERE
				topic_id = ?				
	`
	res, err := m.DB.Query(stmt, id)

	var words []*models.Word

	for res.Next() {
		word := &models.Word{}
		err = res.Scan(&word.ID, &word.Word, &word.Translation)
		if err != nil{
			panic(err)
		}

		words = append(words, word)
	}

	TC := &models.TopicContent{
		TopicID : topic.ID,
		TopicName : topic.Name,
		Words : words,
	}


	return TC, err
}

func (m *TopicModel) AddWordToBase (topicId int, word string, translation string) (int, error){
	stmt := `INSERT INTO words (
				topic_id,
				word,
				translation)
			VALUES (
			?,
			?,
			?)
	`

	result, err := m.DB.Exec(stmt, topicId, word, translation)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err !=nil {
		return 0, err
	}

	return int(id), nil

}

func (m *TopicModel) GetTopic (id int) (*models.Topic, error) {
	stmt := `SELECT
				id,
				name
			FROM
				topic
			WHERE
				id = ?`
	row := m.DB.QueryRow(stmt, id)

	topic := &models.Topic{}
	err := row.Scan(&topic.ID, &topic.Name)
	if err != nil {
		panic(err)
	}

	return topic, err

}

func (m *TopicModel) UpdateTopic (id int, name string) (int, error) {
	stmt := `UPDATE
				topic
			SET
				name = ?
			WHERE
				id = ?`
	
	_, err := m.DB.Exec(stmt, name, id)
	if err != nil{
		return 0, err
	}

	return int(id), nil
} 
