# Exercício Cotação MBA Go Expert

Este projeto demonstra uma aplicação em Go para consultar e armazenar cotações de câmbio USD-BRL usando Clean Architecture.

## Como Executar

### Servidor
- Navegue até `cmd/server`.
- Execute: `go run main.go`.
- O servidor iniciará na porta 8080 e criará o banco de dados SQLite (`cambio.db`) na pasta `cmd/server`.

### Cliente
- Navegue até `cmd/client`.
- Execute: `go run main.go`.
- O cliente consultará a API e salvará a cotação no arquivo `cotacao.txt` na pasta `cmd/client`.

## Estrutura do Projeto
- Usa Clean Architecture com camadas de entidade, caso de uso, interface e infraestrutura.
- Banco de dados: SQLite.
- API externa: AwesomeAPI para cotações.

## Dependências
- Instale: `go get github.com/mattn/go-sqlite3`.
- Execute `go mod tidy` para atualizar módulos.