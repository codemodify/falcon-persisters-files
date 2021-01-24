package persisters

import (
	"github.com/codemodify/falcon/data"
	"github.com/codemodify/falcon/persisters"
)

type filesPersister struct {
}

func NewFilesPersister(folder string) persisters.StorePersister {
	return &filesPersister{}
}

func (thisRef filesPersister) Load() (data.Store, error) {
	return nil
}

func (thisRef filesPersister) Save() error {
	return nil
}
