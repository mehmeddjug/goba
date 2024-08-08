package database

import (
	"errors"
	"sync"

	"github.com/mehmeddjug/goba/content/internal/core"
)

type ContentRepositoryMemory struct {
	content map[string]*core.Content
	mu      sync.Mutex
}

func NewContentRepositoryMemory() *ContentRepositoryMemory {
	return &ContentRepositoryMemory{
		content: make(map[string]*core.Content),
	}
}

func (r *ContentRepositoryMemory) Create(story *core.Content) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.content[story.ID]; exists {
		return errors.New("story already exists")
	}
	r.content[story.ID] = story
	return nil
}

func (r *ContentRepositoryMemory) Get(id string) (*core.Content, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	story, exists := r.content[id]
	if !exists {
		return nil, errors.New("story not found")
	}
	return story, nil
}

func (r *ContentRepositoryMemory) Update(story *core.Content) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.content[story.ID]; !exists {
		return errors.New("story not found")
	}
	r.content[story.ID] = story
	return nil
}

func (r *ContentRepositoryMemory) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.content[id]; !exists {
		return errors.New("user not found")
	}
	delete(r.content, id)
	return nil
}

func (r *ContentRepositoryMemory) GetAll() ([]*core.Content, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var contentList []*core.Content
	for _, story := range r.content {
		contentList = append(contentList, story)
	}
	return contentList, nil
}
