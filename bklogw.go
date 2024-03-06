package bklogw

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zestx5/bklogw/internal/storage"
)

func mkConfDir() {
	dn, err := os.UserHomeDir()
	if err != nil {
		fmt.Print(err)
	}
	err = os.Mkdir(dn+"/.bklogw", 0o744)
	if err != nil {
		fmt.Print(err)
	}
}

func Main() int {
	mkConfDir()
	dn, err := os.UserHomeDir()
	if err != nil {
		fmt.Print(err, "\n")
		return 1
	}
	s, err := storage.Open("sqlite3", dn+"/.bklogw/store.db")

	if err != nil {
		fmt.Print(err, "\n")
		return 1
	}

	if err := s.Close(); err != nil {
		fmt.Print(err, "\n")
		return 1
	}
	return 0
}
