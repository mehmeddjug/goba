package ports

import "github.com/mehmeddjug/goba/users/internal/core"

type ActivityLogRepository interface {
	LogActivity(log core.ActivityLog) error
}
