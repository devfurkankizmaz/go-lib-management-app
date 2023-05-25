# Go Book Management Rest API

This repository demonstrates how to manage books using Golang Rest API, with user authentication and authorization.
It showcases the integration with a database and illustrates how CRUD operations are performed. 
This repository is still in the process of being enhanced and once completed,
the front-end part will also be developed and added to the repository in the near future.

Application demonstrates:

- `REST API` using [Echo](https://echo.labstack.com/)
- `PostgreSQL integration` using [GORM](https://gorm.io)

## Quick Start

Make sure you have Docker installed...
Create and customize your config.yaml file in the root directory.

1. Run `make dev`
- `Its starts postgresql & pgadmin`
2. Run `make migration`
- `Run the migration to prepare the database.`
3. Run `make app`
- `It run the server at`(localhost:8080)

