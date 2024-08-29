# Project Golang Task CRUD

this project explore golang with library atlas and air

- [Atlas for schema databased and migrate version](https://atlasgo.io/)
- [Air - Live reload for go apps](https://github.com/air-verse/air)

## Run project

- before running project
  `go mod tidy`

- project running
  `go run main.go`

- code migrate in database
  `atlas migrate apply --env gorm -u "mysql://root:pass@localhost:3306/db_crud" `

- Generate file migrations
  `atlas migrate diff --env gorm `

Thank you for visiting this Git repository
