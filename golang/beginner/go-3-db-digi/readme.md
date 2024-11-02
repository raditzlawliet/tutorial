# Tutorial - Golang Connect to Database (with Gin & GORM)

We will create an API and connect to Database with (Gin and GORM)

<p align="center">
    <img src="docs/overview.png">
    <img src="docs/result.png" height="350px">
</p>

## What's covered

- REST API with Gin
- Read JSON Request to Model/Struct and do JSON Response
- Connect to Postgres Database with GORM
- Do Query CRUD when calling REST API

### Database

- auths
  - auth_id: bigint, autoincrement
  - account_id: bigint
  - username: varchar
  - password: varchar
- accounts
  - account_id: bigint
  - name: varchar

### API

- /auth/login
- /auth/upsert
- /account/create
- /account/read
- /account/update
- /account/delete
- /account/list
- /account/my

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
