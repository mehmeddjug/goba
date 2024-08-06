package database

import (
	"sync"

	"github.com/mehmeddjug/goba/users/internal/core"
)

type InMemoryActivityLogRepository struct {
	logs []core.ActivityLog
	mu   sync.Mutex
}

func NewInMemoryActivityLogRepository() *InMemoryActivityLogRepository {
	return &InMemoryActivityLogRepository{}
}

func (repo *InMemoryActivityLogRepository) LogActivity(log core.ActivityLog) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.logs = append(repo.logs, log)
	return nil
}
