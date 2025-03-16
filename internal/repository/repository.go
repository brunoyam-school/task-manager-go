package repository

import (
	"awesomeProject5/internal/entity"
	"database/sql"
	"github.com/spf13/viper"
	_ "modernc.org/sqlite"
)

type Repo struct {
	db *sql.DB
}

func New() *Repo {
	db, err := sql.Open("sqlite", viper.GetString("db_path"))
	if err != nil {
		panic(err)
	}

	initQuery := `CREATE TABLE IF NOT EXISTS tasks (
	id TEXT PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT,
	status BOOLEAN,
	priority INTEGER);`

	_, err = db.Exec(initQuery)
	if err != nil {
		panic(err)
	}

	return &Repo{db: db}
}

func (r *Repo) Create(task entity.Task) error {
	query := `INSERT INTO tasks (ID, TITLE, DESCRIPTION, STATUS, PRIORITY) VALUES (?, ?, ?, ?, ?);`

	_, err := r.db.Exec(query, task.ID, task.Title, task.Description, task.Status, task.Priority)
	if err != nil {
		return err
	}

	return nil
}
