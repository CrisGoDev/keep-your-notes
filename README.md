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

and the database is deployed in aws

## Notes

### GET/Notes

Get list of "Notes".

Get Parameters supports:
- *limit* and *offset* get parameters. See
heading "Pagination" above for details.
- *title*
- *body*
- *created_at*
- *updated_at*

Response: 200

```json
{
    "status":200,
    "data":[

    ]

}
```


### GET/Notes/:id

Get the "Notes" info.

Response: 200

```json
{
    "status":200,
    "data":{
        
    }

}
```

### POST/Notes
Create a new note account.

Request:

```json
{
    "title":"Exaple Title",
    "body":"Example body notes"
}
```
| Parameter | Description |Is required|Maximum length|
| ----------- | ----------- |-----------|-----------|
| *title* | the title of the notes | true|100|
| *body* | the body of the notes | true|500|

*Response: 201*

```json
{
    "status":201,
    "data":
    {
    "id":"2c1003cb-dd3a-4dd5-a498-681e7ef5fe33",
    "title":"Exaple Title",
    "body":"Example body notes",
    "created_at":"2021/12/19"
    }
}
```

### PATCH/Notes/:id
Update notes.
 At least you must specifie a field

Request:

```json
{
    "title":"Exaple Title",
    "body":"Example body notes"
}
```
| Parameter | Description |Is required|Maximum length|
| ----------- | ----------- |-----------|-----------|
| *title* | the title of the notes | false|100|
| *body* | the body of the notes | false|500|

*Response: 200*

```json
{
    "status":200,
    "data":
    {
    "id":"2c1003cb-dd3a-4dd5-a498-681e7ef5fe33",
    "title":"Exaple Title",
    "body":"Example body notes",
    "created_at":"2021/12/19",
    "updated_at":"2021/12/19"
    }
}
```

### DELETE /Notes/:Id

Completely delete notes. This cannot be
undone.

Response: 200