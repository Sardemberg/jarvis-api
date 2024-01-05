package commands

type Command interface {
	Execute(params ...string) (string, error)
	GetDescription() string
}
