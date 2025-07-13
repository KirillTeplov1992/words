package store

import "words/internal/models"

type WordRepository struct{
	store *Store
}

func (wr *WordRepository) GetWord (id int) (*models.Word, error){
	stmt := `SELECT
				id,
				word,
				translation
			FROM
				words
			WHERE
				id = ?`
	row := wr.store.db.QueryRow(stmt, id)

	word := &models.Word{}
	err := row.Scan(&word.ID, &word.Word, &word.Translation)
	if err != nil {
		panic(err)
	}

	return word, err
}

func (wr *WordRepository) UpdateWord (id int, word string, translation string) (int, error) {
	stmt := `UPDATE
				 words
			SET
				word = ?,
				translation = ?
			WHERE
				id = ?`

	_, err := wr.store.db.Exec(stmt, word, translation, id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}