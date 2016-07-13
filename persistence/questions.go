package persistence

import (
	"database/sql"
	"github.com/DennisVis/bpt-go/models"
)

type QuestionsDAO struct {
	AppDB DB
}

func (qDao QuestionsDAO) fetchQuestions(query string, args ...interface{}) ([]models.Model, error) {

	rows, err := qDao.AppDB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	questionsMap := make(map[int]models.Model, 1)

	defer rows.Close()

	for rows.Next() {

		var (
			id int
			name string
			language string
			value string
		)

		err := rows.Scan(&id, &name, &language, &value)
		if err != nil {
			return nil, err
		}

		if m, exists := questionsMap[id]; exists {
			if q, ok := m.(models.Question); ok {
				models.Question(q).Labels[language] = value
				questionsMap[id] = q
			}
		} else {
			questionsMap[id] = models.Question{id, name, map[string]string{language: value}}
		}
	}

	return mapToSlice(questionsMap), nil
}

func (qDao QuestionsDAO) insertLabels(tx *sql.Tx, questionId int, labels map[string]string) error {

	for language, value := range labels {

		var labelId int
		err := tx.QueryRow(
			"INSERT INTO labels(language, value) VALUES($1, $2) RETURNING id;",
			language, value).Scan(&labelId)
		if err != nil {
			tx.Rollback()
			return err
		}

		stmt, err := tx.Prepare("INSERT INTO labels_per_question(question_id, label_id) VALUES($1, $2);")
		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = stmt.Exec(questionId, labelId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}

func (qDao QuestionsDAO) All() ([]models.Model, error) {

	return qDao.fetchQuestions(`
		SELECT questions.id, questions.name, labels.language, labels.value
		FROM questions
		LEFT OUTER JOIN labels_per_question
		ON labels_per_question.question_id = questions.id
		LEFT OUTER JOIN labels
		ON labels_per_question.label_id = labels.id;
	`)
}

func (qDao QuestionsDAO) Create(model models.Model) (int, error) {

	question := model.(models.Question)

	tx, err := qDao.AppDB.Begin()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var questionId int
	err = tx.QueryRow("INSERT INTO questions(name) VALUES($1) RETURNING id;", question.Name).Scan(&questionId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = qDao.insertLabels(tx, questionId, question.Labels)

	if err != nil {
		tx.Rollback()
		return 0, err
	} else {
		tx.Commit()
		return questionId, nil
	}
}

func (qDao QuestionsDAO) Read(questionId int) (*models.Model, error) {

	questions, err := qDao.fetchQuestions(`
		SELECT questions.id, questions.name, labels.language, labels.value
		FROM questions
		LEFT OUTER JOIN labels_per_question
		ON labels_per_question.question_id = questions.id
		LEFT OUTER JOIN labels
		ON labels_per_question.label_id = labels.id
		WHERE questions.id = $1;
	`, questionId)

	if err != nil {
		return nil, err
	} else if len(questions) < 1 {
		return nil, nil
	} else {
		return &questions[0], nil
	}
}

func (qDao QuestionsDAO) Update(questionId int, model models.Model) (models.Model, error) {

	question := model.(models.Question)

	tx, err := qDao.AppDB.Begin()
	if err != nil {
		tx.Rollback()
		return question, err
	}

	// Update question
	stmt, err := tx.Prepare("UPDATE questions SET name = $1 WHERE id = $2;")
	if err != nil {
		tx.Rollback()
		return question, err
	}
	_, err = stmt.Exec(question.Name, question.Id)
	if err != nil {
		tx.Rollback()
		return question, err
	}
	//

	// Remove labels
	_, err = tx.Exec("DELETE FROM labels_per_question WHERE question_id = $1;", question.Id)
	if err != nil {
		tx.Rollback()
		return question, err
	}
	//

	// Update labels
	err = qDao.insertLabels(tx, questionId, question.Labels)
	//

	if err != nil {
		tx.Rollback()
		return question, err
	} else {
		tx.Commit()
		return question, nil
	}
}

func (qDao QuestionsDAO) Delete(questionId int) (int, error) {

	res, err := qDao.AppDB.Exec("DELETE FROM questions WHERE id = $1;", questionId)
	ra, err := res.RowsAffected()

	if err != nil {
		return 0, err
	} else {
		return int(ra), err
	}
}
