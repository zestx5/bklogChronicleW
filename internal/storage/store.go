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
	Add(Game) error
	Get(int) (Game, error)
	GetAll() ([]Game, error)
	Delete(int) error
	Update(int, Game) (Game, error)
	Close() error
}

type Store struct {
	DB *sql.DB
}

type Game struct {
	Id    int
	Title string
}

func (s *Store) Add(g Game) error {
	stmt, err := s.DB.Prepare("INSERT INTO backlog(title) VALUES(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(g.Title); err != nil {
		return err
	}
	return nil
}

func (s *Store) Get(id int) (Game, error) {
	stmt, err := s.DB.Prepare("SELECT * FROM backlog WHERE id = ?")
	if err != nil {
		return Game{}, err
	}
	defer stmt.Close()
	var g Game
	err = stmt.QueryRow(id).Scan(&g.Id, &g.Title)
	if err != nil {
		return g, err
	}
	return g, nil

}

func (s *Store) GetAll() ([]Game, error) {
	gs := []Game{}
	q := "SELECT * FROM backlog"
	rows, err := s.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var g Game
		if err := rows.Scan(&g.Id, &g.Title); err != nil {
			return nil, err
		}
		gs = append(gs, g)
	}
	return gs, nil
}

func (s *Store) Delete(id int) error {
	stmt, err := s.DB.Prepare("DELETE FROM backlog WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Update(id int, updatedGame Game) (Game, error) {
	return Game{}, nil
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
