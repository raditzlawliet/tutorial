# Tutorial - Golang Communicating between services with gRPC

Continue from previous material [go-3-db-digi](./../go-3-db-digi/), Now we will decoupled our Auth into single service to handle auth.

While other will be Account Service also act as API Gateway.

One big major point is, auths table now owned by Auth Service (logically it can separate into differentF database)

<p align="center">
    <img src="docs/overview.png">
    <img src="docs/result.png" height="350px">
</p>

## What's covered

- Decoupled into 2 Services: Account and Auth
- Protobuf
- Implement gRPC Handler
- Communicate Account Service to Auth Service via gRPC

### Also it's continue from previous material

- REST API with Gin
- Read JSON Request to Model/Struct and do JSON Response
- Connect to Postgres Database with GORM
- Do Query CRUD when calling REST API

### Database

- auths (Owned by Auth Services)

  - auth_id: bigint, autoincrement
  - account_id: bigint
  - username: varchar
  - password: varchar

- accounts (Owned by Account Services)
  - account_id: bigint
  - name: varchar

### API (Account Service)

- /auth/login -> Auth Service Auth/Login
- /auth/upsert -> Auth Service Auth/Upsert
- /account/create
- /account/read
- /account/update
- /account/delete
- /account/list
- /account/my -> Middleware Validate Token to Auth Service Auth/Validate

### gRPC (Auth Service)

- Auth/Login
- Auth/Upsert
- Auth/Validate

## Not mentioned

- Including Code Tour
  - Each topic has spesific commit
  - Script to automatically reclone with spesific commit, see on [vscode-open-commit.sh](./script/vscode-open-commit.sh)

## How to use

- Clone repository
- Open with VSCode
- Don't forget to install Golang
- Have Postgres and Import DDL to your Database
- Enable Golang VSCode extension
- Enable Codetour VSCode extension (https://marketplace.visualstudio.com/items?itemName=vsls-contrib.codetour)
- Follow Codetour topic and step, each topic has script to automatically reclone with spesific commit.

## Support

If you have any questions or encounter any issues, feel free to open an issue and we will assist you in resolving them.

## Contribute

Feel free to contribute, don't forget to mention if needed

## Author and Contributor

- [@raditzlawliet](https://github.com/raditzlawliet/)
