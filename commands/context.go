package commands

type Context struct {
	Command Command
}

func (c *Context) ExecuteCommand(params ...string) (string, error) {
	return c.Command.Execute(params...)
}
