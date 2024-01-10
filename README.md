# Jarvis API

Jarvis API é um sistema para manipulação de mensagens do WhatsApp Business API, oferecendo funcionalidades para processar comandos e interações.

## Funcionalidades

- **Comandos Personalizados:** Suporte para execução de comandos personalizados enviados através do WhatsApp Business API.
- **Log de Erros:** Registra logs detalhados em caso de falhas na manipulação de payloads.

## Estrutura do Projeto

O projeto está estruturado da seguinte forma:

- `controllers/`: Contém os controladores que lidam com as solicitações HTTP.
- `commands/`: Módulos que implementam comandos específicos.
- `database/`: Configuração e modelos relacionados ao banco de dados.
- `models/`: Definições de modelos de dados.
- `services/`: Serviços responsáveis por funcionalidades específicas.

## Dependências

- [Gin](https://github.com/gin-gonic/gin): Framework web para Go.
- [GORM](https://gorm.io/): ORM para Go.
- Outras dependências especificadas no arquivo go.mod.

## Configuração

Certifique-se de configurar as variáveis de ambiente necessárias, como `WHATSAPP_TOKEN` e qualquer outra variável utilizada no projeto.

## Como Executar

1. Clone o repositório: `git clone https://github.com/seu-usuario/jarvis-api.git`
2. Navegue até o diretório do projeto: `cd jarvis-api`
3. Build a imagem do docker com o comando: `docker build -t jarvis .`
4. Execute o docker-compose do projeto com : `docker-compose up`

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir problemas ou enviar pull requests para melhorar este projeto.

## Licença

Este projeto é licenciado sob a [Licença MIT](LICENSE).

