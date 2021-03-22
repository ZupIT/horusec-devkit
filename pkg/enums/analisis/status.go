package analisis

const (
	Running = "running"
	Success = "success"
	Error   = "error"
)

func Values() []string {
	return []string{
		Running,
		Success,
		Error,
	}
}
