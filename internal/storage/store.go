package storage

import (
	"database/sql"
)

const CreateStr string = `
CREATE TABLE IF NOT EXISTS backlog(
	id INTEGER NOT NULL PRIMARY KEY,
	title TEXT
);
`

type Storer interface {
	Get(int) (Game, error)
	GetAll() []Game
	Delete(int) bool
	Update(int, Game) Game
	Close() error
}

type Store struct {
	DB *sql.DB
}

type Game struct {
	Name string
}

func (s *Store) Get(id int) (Game, error) {
	return Game{}, nil
}

func (s *Store) GetAll() []Game {
	return nil
}

func (s *Store) Delete(id int) bool {
	return true
}

func (s *Store) Update(id int, updatedGame Game) Game {
	return Game{}
}

func (s *Store) Close() error {
	if err := s.DB.Close(); err != nil {
		return err
	}
	return nil
}

func Open(driver string, name string) (Storer, error) {
	db, err := sql.Open(driver, name)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	if _, err := db.Exec(CreateStr); err != nil {
		return nil, err
	}

	s := &Store{
		db,
	}
	return s, nil
}
