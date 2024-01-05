package commands

type HelloCommand struct{}

func (h *HelloCommand) Execute() (string, error) {
	return "Olá, senhor. Em que posso te ajudar hoje?", nil
}
