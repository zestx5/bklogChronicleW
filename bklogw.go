package bklogw

import (
	"fmt"

	"github.com/zestx5/bklogw/internal/storage"
)

func Main() int {
	s, err := storage.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Print(err)
		return 1
	}
	s.Close()
	return 0
}
