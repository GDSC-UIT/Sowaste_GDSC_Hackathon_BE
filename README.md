# Sowaste GDSC Hackathon BE

## Framework: Gin

Root API URL: https://sowaste.up.railway.app

**Note**: No authentication requires

## API Routes

<hr />

- `/api/dictionaries` : `GET` all dictionaries or `POST` to create a dictionary
- `/api/dictionaries/:id` : `GET, PUT, DELETE` a specific dictionary
- `/api/articles`: `GET` all articles or `POST` to create an article
- `/api/articles/:id` : `GET, PUT, DELETE` a specific article
- `/api/questions` : `GET` all questions or `POST` to create a question
- `/api/questions/:id` : `GET, PUT, DELETE` a specific question
- `/api/options` : `GET` all options or `POST` to create an option
- `/api/options/:id` : `GET, PUT, DELETE` a specific option

<hr />

## Folder structure

`go` - root folder
  - `api`
  - `internal`
    - `app` - The point where all our dependencies and logic are collected and run the app. The run method that is called from /cmd.
    - `config` - Initialization of the general app configurations that we wrote in the root of the project.
    - `database` - The files contain methods for interacting with databases.
    - `models` - The structures of database tables.
    - `services` - The entire business logic of the application.
    - `transport` - Here we store http-server settings, handlers, ports, etc.
    - `utils` - Utilities for the app
  - `go.mod`
  - `README.md`
