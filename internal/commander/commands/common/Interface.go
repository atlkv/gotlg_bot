package commands

type CMDExecutor interface {
	Execute() (string, error)
	GetPath() string
	Info() string
}
