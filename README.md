Etapas efetuadas:
1. Criado projeto GO (`go mod init project_name`)
2. Definido especificações da API no arquivo [internal/api/spec/journey.spec.json](https://github.com/jeffersoncbd/nlw-2024-journey-app/blob/main/internal/api/spec/journey.spec.json) (OpenAPI 3.0.0).
3. Utilizado o pacote [goapi-gen](https://github.com/discord-gophers/goapi-gen) para gerar os boilerplates (interfaces) que atendem as especificações acima.
4. Instalada e atualizada todas as dependencias necessárias utilizando `go mod tidy` e `go get -u ./...`
5. Criado [docker-compose-local.yml](https://github.com/jeffersoncbd/nlw-2024-journey-app/blob/main/docker-compose-local.yml) para criar um container PostgreSQL.
6. Utilizado o pacote [tern](https://github.com/jackc/tern) para criar e executar as [migrations](https://github.com/jeffersoncbd/nlw-2024-journey-app/tree/main/internal/pgstore/migrations) no banco de dados.
7. Implementado as [queries](https://github.com/jeffersoncbd/nlw-2024-journey-app/blob/main/internal/pgstore/queries/queries.sql) do projeto.
8. Utilizado o pacote [sqlc](https://github.com/sqlc-dev/sqlc) para compilar as queries SQL acima em funções GO que às executam.
