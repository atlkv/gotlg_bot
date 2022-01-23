package commands

type CMDExecuter interface {
	Execute() (string, error)
	GetPath() string
	Info() string	
}
