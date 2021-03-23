package analysis

type Status string

const (
	Running Status = "running"
	Success Status = "success"
	Error   Status = "error"
)

func Values() []Status {
	return []Status{
		Running,
		Success,
		Error,
	}
}

func (s Status) ToString() string {
	return string(s)
}
