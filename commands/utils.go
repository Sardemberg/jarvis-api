package commands

import (
	"errors"
	"fmt"
	"regexp"
)

func TransformParams(params ...string) (map[string]string, error) {
	mapperParams := make(map[string]string)

	for _, param := range params {
		re := regexp.MustCompile(`--([^=]+)=(.+)`)

		matches := re.FindStringSubmatch(param)

		if len(matches) != 3 {
			return nil, errors.New(fmt.Sprintf("Formato da string '%s' inválido", param))
		}

		key := matches[1]
		value := matches[2]

		mapperParams[key] = value
	}

	return mapperParams, nil
}
