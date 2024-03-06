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
	t.Parallel()
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
}

func TestGetReturnsErrorWhenNoGame(t *testing.T) {
	t.Parallel()
	_, err := db.Get(5)
	if err == nil {
		t.Error("want error when no game, got nothing")
	}
}
