package option

import (
	"container/list"
	"sync"
)

var (
	repo *list.List
	lock = &sync.Mutex{}
)

func GetRepo() *list.List {
	if repo == nil {
		lock.Lock()
		defer lock.Unlock()

		if repo == nil {
			repo = list.New()
		}
	}

	return repo
}
