package commands

type Context struct {
	Command Command
}

func (c *Context) ExecuteCommand() (string, error) {
	return c.Command.Execute()
}
