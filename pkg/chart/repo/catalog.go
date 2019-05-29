package repo

import (
	"fmt"
	"net/url"
	"sync"
)

type CacheFactory func(name string) (Cache, error)

type Catalog struct {
	factory CacheFactory
	repos   map[string]*Repo
	sync.Mutex
}

func NewCatalog(factory CacheFactory) *Catalog {
	return &Catalog{
		repos:   make(map[string]*Repo),
		factory: factory,
	}
}

func (c *Catalog) CreateRepoIfNotExist(repoURL string) (*Repo, error) {
	if _, err := url.ParseRequestURI(repoURL); err != nil {
		return nil, err
	}

	c.Lock()
	defer c.Unlock()

	name := url2name(repoURL)
	repo, ok := c.repos[name]
	if !ok {
		cache, err := c.factory(name)
		if err != nil {
			return nil, fmt.Errorf("failed to create cache: %v", err)
		}

		repo = NewRepo(repoURL, cache)

		c.repos[name] = repo
	}

	return repo, nil
}
