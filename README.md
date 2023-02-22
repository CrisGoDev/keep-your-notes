# keep-your-notes
This is an application focused on being able to manage notes that will be stored in a database. the scope of this is: Add notes Delete notes Update notes Search notes by name, content and date.

## Requirements to start the project
- Docker
- Go

# How to start the project locally?
1. Open the terminal
2. Execute `docker pull mysql`
3. Execute `docker compose build`
4. Execute `docker compose up -d`
5. please create a file **.env** and use the variable specified on **.env.example**
6. Execute ```go run main.go```
(Note)
### if you are running the api on {{ server }}
**when making a request to an endpoint for the first time, please wait a moment.**

### if you are having problem to run locally 
please try to create the build again and execute `docker compose up -d --remove-orphans`
and try `go run main.go`
now you can use the endpoints documented, with the base url ```http://localhost:8000```

# API Documentation V1

This file documents the API implemented at
https://github.com/CrisGoDev/keep-your-notes , which is a
REST api that can access one of the following objects:
- Notes - A User can create a Notes

# Pagination
The endpoint **get all** has the query parameter

| Parameter | Values posible |Default|Description|
| ----------- | ----------- |-----------|-----------|
| *page* | integer numbers | 1|the page you want|
| *limit* | integer numbers | 15|the quantity of records|


## Documentation Postman
 There is a file called *keep-notes.postman_collection* you can import this collection
 in this collection two variables to use **local** and **server** are configured.
 - you can swap between it using {{local}} and {{server}} by default it is {{server}}.
 on Postman ![Postman](https://miro.medium.com/v2/resize:fit:1400/format:webp/1*_Hc7qFt6iJqNGq1RmnCiMg@2x.png "postman")

## Unite test
in the project are configure some unit test to execute this, you can run `go test -v .\pkg\meta\`

## Endpoints

All endpoints documented below are relative to the following base urls:

Development: https://keep-notes-xx4n.onrender.com

and the database is deployed in aws
![Aws rds](https://d1by4p17n947rt.cloudfront.net/icon/1d374ed2a6bcf601d7bfd4fc3dfd3b5d-c9f69416d978016b3191175f35e59226.svg "aws rds")

## Notes

### GET/notes

Get list of "Notes".

Get Parameters supports:
- *limit* and *paget* get parameters. See
heading "Pagination" above for details.
- *title*
- *body*
- *order*
    - the posibles values are *asc* & *desc*
- *order_by* 
    - the posibles values are *title* ,*body* & *created_at*
Response: 200

```json
{
    "status": 200,
    "data": [
        {
            "id": "25da7aa4-0d99-4e98-be4c-ef1a3f8a56da",
            "title": "Updated Again",
            "body": "Updated again",
            "created_at": "2023-02-20T14:09:47.086-06:00",
            "updated_at": "2023-02-20T20:35:12.042-06:00"
        },
        {
            "id": "f217d9f3-400c-41fd-ab48-3d83c219b409",
            "title": "pruebasasas",
            "body": "spe body notesjfgsdhgjhsd",
            "created_at": "2023-02-20T14:06:26.206-06:00",
            "updated_at": "2023-02-20T14:06:26.206-06:00"
        },
        {
            "id": "c29a363e-d9a1-4603-a586-837622eff4cb",
            "title": "aadasfa",
            "body": "spe body notesjfgsdhgjhsd",
            "created_at": "2023-02-22T04:24:29.726-06:00",
            "updated_at": "2023-02-22T04:24:29.726-06:00"
        },
        {
            "id": "537ae82f-1b45-4708-8384-4180a5d72534",
            "title": "Delete * from where = true",
            "body": "spe body notesjfgsdhgjhsd",
            "created_at": "2023-02-20T14:06:34.242-06:00",
            "updated_at": "2023-02-20T14:06:34.242-06:00"
        },
        {
            "id": "ab589771-5cc4-48cc-a7f8-ff2ff1b0912c",
            "title": "Exaple Title",
            "body": "Example body notes",
            "created_at": "2023-02-19T23:34:38.635-06:00",
            "updated_at": "2023-02-19T23:34:38.635-06:00"
        },
        {
            "id": "a405faa6-354c-4393-9006-45c94deac9a5",
            "title": "Exaple Title",
            "body": "Example body notes",
            "created_at": "2023-02-19T23:35:03.474-06:00",
            "updated_at": "2023-02-19T23:35:03.474-06:00"
        },
        {
            "id": "b9f4f231-a542-4605-b11b-c4aa5b724aa0",
            "title": "specific",
            "body": "Example body notes",
            "created_at": "2023-02-19T23:40:21.976-06:00",
            "updated_at": "2023-02-19T23:40:21.976-06:00"
        },
        {
            "id": "c4e54be4-1917-4220-a7b3-27c31b8cd873",
            "title": "Exaple Title",
            "body": "Example body notes",
            "created_at": "2023-02-19T23:34:58.515-06:00",
            "updated_at": "2023-02-19T23:34:58.515-06:00"
        },
        {
            "id": "c86e1d30-45fe-43e7-a6b9-c10d742dae54",
            "title": "Exaple Title",
            "body": "Example body notes",
            "created_at": "2023-02-19T23:34:30.564-06:00",
            "updated_at": "2023-02-19T23:34:30.564-06:00"
        },
        {
            "id": "dd295309-7683-4035-ad41-75ba99b37bd5",
            "title": "Exaple Title",
            "body": "Example body notes",
            "created_at": "2023-02-19T23:34:28.072-06:00",
            "updated_at": "2023-02-19T23:34:28.072-06:00"
        }
    ],
    "meta": {
        "total_count": 26,
        "per_page": 10,
        "page_count": 3,
        "page": 1
    }
}
```

- **total_count** : the quantity of records
- **per_page** : the quantity of records per pages
- **page_count** : the quantity of pages
- **page** : the current page you are

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