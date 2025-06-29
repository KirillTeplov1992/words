package mysql

import (
	"database/sql"
	"words/pkg/models"
)

type WordModel struct{
	DB *sql.DB
}

func (mw *WordModel) GetWord (id int) (*models.Word, error){
	stmt := `SELECT
				id,
				word,
				translation
			FROM
				words
			WHERE
				id = ?`
	row := mw.DB.QueryRow(stmt, id)

	word := &models.Word{}
	err := row.Scan(&word.ID, &word.Word, &word.Translation)
	if err != nil {
		panic(err)
	}

	return word, err
}

func (mw *WordModel) UpdateWord (id int, word string, translation string) (int, error) {
	stmt := `UPDATE
				 words
			SET
				word = ?,
				translation = ?
			WHERE
				id = ?`

	_, err := mw.DB.Exec(stmt, word, translation, id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}