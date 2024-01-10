package commands

import (
	"fmt"
	"jarvisapi/services"
	"strings"
)

type NewsCommand struct{}

func (n *NewsCommand) Execute(params ...string) (string, error) {
	articles, err := services.GetNews()

	if err != nil {
		return "", err
	}

	mappedParams, err := TransformParams(params...)

	if err != nil {
		return "", err
	}

	for _, article := range articles {
		message := fmt.Sprintf(`
				*Título*: %s
				*Fonte*: %s
				*Url*: %s 
			`,
			article.Title,
			article.Source.Name,
			article.Url,
		)

		message = removeWhiteLines(message)

		services.SendNewMessage(message, mappedParams["number"])
	}

	return "Essas são as principais notícias do dia!", nil
}

func removeWhiteLines(message string) string {
	messageSplitted := strings.Split(message, "\n")

	for i, line := range messageSplitted {
		messageSplitted[i] = strings.TrimSpace(line)
	}

	return strings.Join(messageSplitted, "\n\n")
}

func (n *NewsCommand) GetDescription() string {
	return "Apenas notícias por aqui..."
}
