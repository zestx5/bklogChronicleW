package storage_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zestx5/bklogw/internal/storage"
)

var db storage.Storer

func setup() {
	memDb, _ := sql.Open("sqlite3", ":memory:")

	memDb.Exec(storage.CreateStr)

	db = &storage.Store{DB: memDb}
}

func teardown() {
	db.Close()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestStoreImplementsStorer(t *testing.T) {
	t.Parallel()
	var _ storage.Storer = &storage.Store{}
}

func TestAddGameWorksAsExpected(t *testing.T) {
	want := storage.Game{Id: 1, Title: "Tekken 8"}
	err := db.Add(want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := db.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
	db.Delete(want.Id)
}

func TestGetAllWorksAsExpected(t *testing.T) {
	want := []storage.Game{
		{Id: 1, Title: "Tekken 8"},
		{Id: 2, Title: "Last Epoch"},
		{Id: 3, Title: "Nioh"},
	}
	for _, v := range want {
		err := db.Add(v)
		if err != nil {
			t.Fatal(err)
		}
	}
	got, err := db.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 {
		t.Errorf("want %d entites, got %d", len(want), len(got))
	}
	for _, v := range want {
		db.Delete(v.Id)
	}
}

func TestDeleteWorksAsExpected(t *testing.T) {
	g := storage.Game{Id: 1, Title: "Tekken 8"}
	err := db.Add(g)
	if err != nil {
		t.Fatal(err)
	}

	db.Delete(g.Id)
	got, err := db.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Errorf("want 0 rows in db, got %d", len(got))
	}
}

func TestUpdateWorksAsExpected(t *testing.T) {
	g := storage.Game{Id: 1, Title: "Tekken 8"}
	err := db.Add(g)
	if err != nil {
		t.Fatal(err)
	}

	want := storage.Game{Id: 1, Title: "Persona 5"}
	got, err := db.Update(g.Id, want)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetReturnsErrorWhenNoGame(t *testing.T) {
	_, err := db.Get(-1)
	if err == nil {
		t.Error("want error when no game, got nothing")
	}
}
