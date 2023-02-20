# keep-your-notes
This is an application focused on being able to manage notes that will be stored in a database. the scope of this is: Add notes Delete notes Update notes Search notes by name, content and date.

## Requirements to start the project
- Docker
- Go

# How to start the project locally?
1. Open the terminal
2. Execute `docker compose build`
3. Execute `docker compose up -d`
4. please create a file **.env** and use the variable specified on **.env.example**
5. Execute ```go run main.go```

now you can use the endpoints documented, with the base url ```http://localhost:8000```

# API Documentation V1

This file documents the API implemented at
https://github.com/CrisGoDev/keep-your-notes , which is a
REST api that can access one of the following objects:

- Notes - A User can create a Notes

## Endpoints

All endpoints documented below are relative to the following base urls:

Development: http://keepyournote/api/v1

and the database is in aws