package store

import "words/internal/models"

type TopicRepository struct{
	store *Store
}

func (tr *TopicRepository) GetTopics() ([]*models.Topic) {
	stmt := `SELECT
				id,
				name
			FROM
				topic`
	res, err := tr.store.db.Query(stmt)
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

func (tr *TopicRepository) GetTopic (id int) (*models.Topic, error) {
	stmt := `SELECT
				id,
				name
			FROM
				topic
			WHERE
				id = ?`
	row := tr.store.db.QueryRow(stmt, id)

	topic := &models.Topic{}
	err := row.Scan(&topic.ID, &topic.Name)
	if err != nil {
		panic(err)
	}

	return topic, err

}

func (tr *TopicRepository) Insert (name string) (int, error){
	stmt := `INSERT INTO topic (
				name)
			VALUES (
				?)`
	result, err := tr.store.db.Exec(stmt, name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err !=nil {
		return 0, err
	}

	return int(id), nil
}

func (tr *TopicRepository) GetContentOfTopicById (id int)(*models.TopicContent, error){
	topic, err := tr.GetTopic(id)

	stmt := `SELECT 
				id,
				word,
				translation
			FROM 
				words
			WHERE
				topic_id = ?				
	`
	res, err := tr.store.db.Query(stmt, id)

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

func (tr *TopicRepository) AddWordToBase (topicId int, word string, translation string) (int, error){
	stmt := `INSERT INTO words (
				topic_id,
				word,
				translation)
			VALUES (
			?,
			?,
			?)
	`

	result, err := tr.store.db.Exec(stmt, topicId, word, translation)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err !=nil {
		return 0, err
	}

	return int(id), nil

}

func (tr *TopicRepository) UpdateTopic (id int, name string) (int, error) {
	stmt := `UPDATE
				topic
			SET
				name = ?
			WHERE
				id = ?`
	
	_, err := tr.store.db.Exec(stmt, name, id)
	if err != nil{
		return 0, err
	}

	return int(id), nil
}