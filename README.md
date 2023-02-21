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

## Documentation Postman
 There is a file called *keep-notes.postman_collection* you can import this collection
 in this collection two variables to use **local** and **server** are configured.
 - you can swap between it using {{local}} and {{server}} by default it is {{server}}.
 on Postman ![Postman](https://miro.medium.com/v2/resize:fit:1400/format:webp/1*_Hc7qFt6iJqNGq1RmnCiMg@2x.png "postman")

## Endpoints

All endpoints documented below are relative to the following base urls:

Development: https://keep-notes-xx4n.onrender.com

and the database is deployed in aws
![Aws rds](https://d1by4p17n947rt.cloudfront.net/icon/1d374ed2a6bcf601d7bfd4fc3dfd3b5d-c9f69416d978016b3191175f35e59226.svg "aws rds")

## Notes

### GET/notes

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
    "status": 200,
    "data": [
        {
            "id": "0e363dc4-1135-4be4-8c05-129e927bc717",
            "title": "specific",
            "body": "spe body notes",
            "created_at": "2023-02-19T23:40:54.326-06:00",
            "updated_at": "2023-02-19T23:40:54.326-06:00"
        },
        {
            "id": "b9f4f231-a542-4605-b11b-c4aa5b724aa0",
            "title": "specific",
            "body": "Example body notes",
            "created_at": "2023-02-19T23:40:21.976-06:00",
            "updated_at": "2023-02-19T23:40:21.976-06:00"
        },
        {
            "id": "8f9dc431-a130-4051-8afa-b92559facde3",
            "title": "Exaple Title",
            "body": "Example body notes",
            "created_at": "2023-02-19T23:35:05.985-06:00",
            "updated_at": "2023-02-19T23:35:05.985-06:00"
        }
    ],
    "meta": {
        "total_count": 22,
        "per_page": 3,
        "page_count": 8,
        "page": 1
    }
}
```


### GET/notes/:id

Get the "Notes" info.

Response: 200

```json
{
    "status": 200,
    "data": {
        "id": "0e363dc4-1135-4be4-8c05-129e927bc717",
        "title": "specific",
        "body": "spe body notes"
    }
}
```

### POST/notes
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
    "status": 201,
    "data": {
        "id": "0dbf0064-e593-4716-a27c-1cd293d832bb",
        "title": "Exaple Title",
        "body": "Example body notes"
    }
}
```

### PATCH/notes/:id
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