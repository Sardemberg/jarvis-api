package commands

import "fmt"

type HelloCommand struct{}

func (h *HelloCommand) Execute(params ...string) (string, error) {
	fmt.Println(TransformParams(params...))

	return "Olá, senhor. Em que posso te ajudar hoje?", nil
}

func (h *HelloCommand) GetDescription() string {
	return "Olá, senhor. Eu sou o Jarvis, um assistente virtual criado para serví-lo e lembrá-lo de suas tarefas. Digite /help para mais informações"
}
