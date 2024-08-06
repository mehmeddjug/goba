package core

type ActivityLog struct {
	ID        string
	UserID    string
	Action    string
	Timestamp string
}

type ActivityLogRepository interface {
	LogActivity(log ActivityLog) error
}
