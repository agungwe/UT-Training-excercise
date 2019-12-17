package store

import (
	"fmt"
)

type Store struct {
	BranchStore
	StoreName string
	Owner     string
}

type BranchStore struct {
	StoreName   string
	OwnerBranch string
}

func (s Store) GetStore() {
	fmt.Println(s)
}

func (b BranchStore) GetBranchStore() {
	fmt.Println(b)
}
