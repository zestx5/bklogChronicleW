package storage

import "database/sql"

type Storer interface {
	Get(int) (any, error)
	GetAll() []any
	Delete(int) bool
	Update(int, any) any
}

type Store struct {
	db *sql.DB
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
	if err := s.db.Close(); err != nil {
		return err
	}
	return nil
}

func Open(driver string, name string) (*Store, error) {
	db, err := sql.Open(driver, name)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	s := &Store{
		db,
	}
	return s, nil
}
