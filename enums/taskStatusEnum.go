package enums
type TaskStatus string

const (
    StatusPending   TaskStatus = "pending"
    StatusCompleted TaskStatus = "completed"
)

func (s TaskStatus) IsValid() bool {
    switch s {
    case StatusPending, StatusCompleted:
        return true
    }
    return false
}