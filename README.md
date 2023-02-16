# keep-your-notes
This is an application focused on being able to manage notes that will be stored in a database. the scope of this is: Add notes Delete notes Update notes Search notes by name, content and date.



# API Documentation V1

This file documents the API implemented at
https://github.com/CrisGoDev/keep-your-notes , which is a
REST api that can access one of the following objects:

- User - Someone with access to this API.
- User Session - creating a session is the first step in accessing the API
- Notes - A User can create a Notes, but frst must log

## Endpoints

All endpoints documented below are relative to the following base urls:

Development: http://keepyournote/api/v1

## Authorization

Begin a user session either by creating a new user account with POST /user or by logging in with POST /user/session/. Both provide a JSON object with the key "session". Provide this session via the HTTP Authorization header, eg:

Authorization: Bearer <session_key>

## Pagination

All GET endpoints which list records automatically return up to 100 of the most
recently updated items in order of update time with the most recent first.
Creation counts as an update.

All such endpoints also provide "before" and "after" optional query parameters,
designed for efficiently updating application side caches.

### Empty Cache

Query endpoint with without before or after parameters, getting up to 100
of the most recent items. Save this list for further operations. Ex: 

```
GET <base>/client/
```
